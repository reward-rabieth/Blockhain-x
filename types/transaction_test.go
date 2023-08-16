package types

import (
	"fmt"
	"github.com/reward-rabieth/blockchain/crypto"
	"github.com/reward-rabieth/blockchain/proto"
	"github.com/reward-rabieth/blockchain/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	fromAddress := fromPrivKey.Public().Address().Byte()
	toPrivkey := crypto.GeneratePrivateKey()
	toAddress := toPrivkey.Public().Address().Byte()
	input := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}
	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}
	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}
	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}
	sig := SignTransaction(fromPrivKey, tx)
	input.Signature = sig.Bytes()
	assert.True(t, VerifyTransaction(tx))
	fmt.Printf("%+v\n", tx)
}
