package types

import (
	"encoding/hex"
	"fmt"
	"github.com/reward-rabieth/blockchain/crypto"
	"github.com/reward-rabieth/blockchain/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignBlock(t *testing.T) {
	var (
		block   = util.RandomBlock()
		privKey = crypto.GeneratePrivateKey()
		pubkey  = privKey.Public()
	)
	sig := SignBlock(privKey, block)
	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(pubkey, HashBlock(block)))
}

func TestHashBlock(t *testing.T) {
	blocK := util.RandomBlock()
	hash := HashBlock(blocK)
	fmt.Println(hex.EncodeToString(hash))
	assert.Equal(t, 32, len(hash))
}
