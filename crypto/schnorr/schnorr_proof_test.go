package schnorr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/binance-chain/tss-lib/common/random"
	"github.com/binance-chain/tss-lib/crypto"
	. "github.com/binance-chain/tss-lib/crypto/schnorr"
	"github.com/binance-chain/tss-lib/tss"
)

func TestSchnorrProof(t *testing.T) {
	u := random.GetRandomPositiveInt(tss.EC().Params().N)
	uG := crypto.ScalarBaseMult(tss.EC(), u)
	proof := NewSchnorrProof(u, uG)

	assert.True(t, proof.Alpha.IsOnCurve(tss.EC()))
	assert.NotZero(t, proof.Alpha.X())
	assert.NotZero(t, proof.Alpha.Y())
	assert.NotZero(t, proof.T)
}

func TestSchnorrProofVerify(t *testing.T) {
	u := random.GetRandomPositiveInt(tss.EC().Params().N)
	X := crypto.ScalarBaseMult(tss.EC(), u)

	proof := NewSchnorrProof(u, X)
	res := proof.Verify(X)

	assert.True(t, res, "verify result must be true")
}

func TestSchnorrProofVerifyBad(t *testing.T) {
	u  := random.GetRandomPositiveInt(tss.EC().Params().N)
	u2 := random.GetRandomPositiveInt(tss.EC().Params().N)
	X := crypto.ScalarBaseMult(tss.EC(), u)
	X2 := crypto.ScalarBaseMult(tss.EC(), u2)

	proof := NewSchnorrProof(u2, X2)
	res := proof.Verify(X)

	assert.False(t, res, "verify result must be false")
}
