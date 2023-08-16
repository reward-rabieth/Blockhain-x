package util

import (
	randc "crypto/rand"
	"github.com/reward-rabieth/blockchain/proto"
	"io"
	"math/rand"
	"time"
)

func RandomHash() []byte {
	hash := make([]byte, 32)
	_, err := io.ReadFull(randc.Reader, hash)
	if err != nil {
		return nil
	}
	return hash

}

func RandomBlock() *proto.Block {
	header := &proto.Header{
		Version:   1,
		Height:    int32(rand.Intn(1000)),
		PrevHash:  RandomHash(),
		RootHash:  RandomHash(),
		Timestamp: time.Now().UnixNano(),
	}
	return &proto.Block{
		Header: header,
	}
}
