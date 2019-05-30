package vss_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/binance-chain/tss-lib/common/math"
	"github.com/binance-chain/tss-lib/crypto/vss"
)

func TestCreate(t *testing.T) {
	num, threshold := 3, 2

	secret := math.GetRandomPositiveInt(vss.EC.N)
	t.Log(secret)

	ids := make([]*big.Int, 0)
	for i := 0; i < num; i++ {
		ids = append(ids, math.GetRandomPositiveInt(vss.EC.N))
	}

	params, polyGs, shares, err := vss.Create(threshold, secret, ids)
	assert.Nil(t, err)

	assert.Equal(t, threshold, params.Threshold)
	assert.Equal(t, num, params.NumShares)

	assert.Equal(t, threshold, len(polyGs.PolyG))

	// ensure that each polyGs len() == 2 and non-zero
	for i := range polyGs.PolyG {
		assert.Equal(t, threshold, len(polyGs.PolyG[i]))
		assert.NotZero(t, polyGs.PolyG[i][0])
		assert.NotZero(t, polyGs.PolyG[i][1])
	}

	t.Log(polyGs)
	t.Log(shares)
}

func TestVerify(t *testing.T) {
	num, threshold := 3, 2

	secret := math.GetRandomPositiveInt(vss.EC.N)

	ids := make([]*big.Int, 0)
	for i := 0; i < num; i++ {
		ids = append(ids, math.GetRandomPositiveInt(vss.EC.N))
	}

	_, polyGs, shares, err := vss.Create(threshold, secret, ids)
	assert.NoError(t, err)

	for i := 0; i < num; i++ {
		assert.True(t, shares[i].Verify(polyGs))
	}
}
