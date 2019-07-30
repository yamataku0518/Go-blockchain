package keygen

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/binance-chain/tss-lib/common"
	"github.com/binance-chain/tss-lib/crypto"
	cmt "github.com/binance-chain/tss-lib/crypto/commitments"
	"github.com/binance-chain/tss-lib/crypto/paillier"
	"github.com/binance-chain/tss-lib/crypto/vss"
	"github.com/binance-chain/tss-lib/tss"
)

// Implements Party
// Implements Stringer
var _ tss.Party = (*LocalParty)(nil)
var _ fmt.Stringer = (*LocalParty)(nil)

type (
	LocalParty struct {
		*tss.BaseParty
		params *tss.Parameters

		temp LocalPartyTempData
		data LocalPartySaveData

		// messaging
		end chan<- LocalPartySaveData
	}

	// Everything in LocalPartySaveData is saved locally to user's HD when done
	LocalPartySaveData struct {
		// secret fields (not shared, but stored locally)
		Xi, ShareID *big.Int             // xi, kj
		PaillierSk  *paillier.PrivateKey // ski

		// public keys (Xj = uj*G for each Pj)
		BigXj       []*crypto.ECPoint     // Xj
		ECDSAPub    *crypto.ECPoint       // y
		PaillierPks []*paillier.PublicKey // pkj

		// h1, h2 for range proofs
		NTildej, H1j, H2j []*big.Int

		// original indexes (ki in signing preparation phase)
		Index int // added for unit test
		Ks    []*big.Int
	}

	LocalPartyMessageStore struct {
		// messages
		kgRound1CommitMessages       []*KGRound1CommitMessage
		kgRound2VssMessages          []*KGRound2VssMessage
		kgRound2DeCommitMessages     []*KGRound2DeCommitMessage
		kgRound3PaillierProveMessage []*KGRound3PaillierProveMessage
	}

	LocalPartyTempData struct {
		LocalPartyMessageStore

		// temp data (thrown away after keygen)
		ui            *big.Int // used for tests
		KGCs          []*cmt.HashCommitment
		vs            vss.Vs
		shares        vss.Shares
		deCommitPolyG cmt.HashDeCommitment
	}
)

// Exported, used in `tss` client
func NewLocalParty(
	params *tss.Parameters,
	out chan<- tss.Message,
	end chan<- LocalPartySaveData,
) *LocalParty {
	partyCount := params.PartyCount()
	p := &LocalParty{
		BaseParty: &tss.BaseParty{
			Out: out,
		},
		params: params,
		temp: LocalPartyTempData{},
		data: LocalPartySaveData{Index: params.PartyID().Index},
		end:  end,
	}
	// msgs init
	p.temp.KGCs = make([]*cmt.HashCommitment, partyCount)
	p.temp.kgRound1CommitMessages = make([]*KGRound1CommitMessage, partyCount)
	p.temp.kgRound2VssMessages = make([]*KGRound2VssMessage, partyCount)
	p.temp.kgRound2DeCommitMessages = make([]*KGRound2DeCommitMessage, partyCount)
	p.temp.kgRound3PaillierProveMessage = make([]*KGRound3PaillierProveMessage, partyCount)
	// data init
	p.data.BigXj = make([]*crypto.ECPoint, partyCount)
	p.data.PaillierPks = make([]*paillier.PublicKey, partyCount)
	p.data.NTildej = make([]*big.Int, partyCount)
	p.data.H1j, p.data.H2j = make([]*big.Int, partyCount), make([]*big.Int, partyCount)
	// round init
	round := newRound1(params, &p.data, &p.temp, out)
	p.Round = round
	return p
}

func (p *LocalParty) String() string {
	return fmt.Sprintf("id: %s, round: %d", p.PartyID(), p.Round.RoundNumber())
}

func (p *LocalParty) PartyID() *tss.PartyID {
	return p.params.PartyID()
}

func (p *LocalParty) Start() *tss.Error {
	p.Lock()
	defer p.Unlock()
	if round, ok := p.Round.(*round1); !ok || round == nil {
		return p.WrapError(errors.New("could not start. this party is in an unexpected state. use the constructor and Start()"))
	}
	common.Logger.Infof("party %s: keygen round %d starting", p.Round.Params().PartyID(), 1)
	return p.Round.Start()
}

func (p *LocalParty) Update(msg tss.Message, phase string) (ok bool, err *tss.Error) {
	return tss.BaseUpdate(p, msg, phase)
}

func (p *LocalParty) StoreMessage(msg tss.Message) (bool, *tss.Error) {
	fromPIdx := msg.GetFrom().Index

	// switch/case is necessary to store any messages beyond current round
	// this does not handle message replays. we expect the caller to apply replay and spoofing protection.
	switch msg.(type) {

	case KGRound1CommitMessage: // Round 1 broadcast messages
		r1msg := msg.(KGRound1CommitMessage)
		p.temp.kgRound1CommitMessages[fromPIdx] = &r1msg

	case KGRound2VssMessage: // Round 2 P2P messages
		r2msg1 := msg.(KGRound2VssMessage)
		p.temp.kgRound2VssMessages[fromPIdx] = &r2msg1 // just collect

	case KGRound2DeCommitMessage:
		r2msg2 := msg.(KGRound2DeCommitMessage)
		p.temp.kgRound2DeCommitMessages[fromPIdx] = &r2msg2

	case KGRound3PaillierProveMessage:
		r3msg := msg.(KGRound3PaillierProveMessage)
		p.temp.kgRound3PaillierProveMessage[fromPIdx] = &r3msg

	default: // unrecognised message, just ignore!
		common.Logger.Warningf("unrecognised message ignored: %v", msg)
		return false, nil
	}
	return true, nil
}

func (p *LocalParty) Finish() {
	p.end <- p.data
}
