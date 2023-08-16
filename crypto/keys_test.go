package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), privKeyLen)
	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubkey := privKey.Public()
	msg := []byte("foo bar barz")
	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(pubkey, msg))
	//Test with invalid message
	assert.False(t, sig.Verify(pubkey, []byte("foo")))
	//Test with invalid pubKey
	invalidPrivKey := GeneratePrivateKey()
	invalidPubKey := invalidPrivKey.Public()
	assert.False(t, sig.Verify(invalidPubKey, msg))

}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed        = "4d191e9ace1b333bf163131548a5c4bb57b7b6052abf07e09f79a7f8487d1756"
		privKey     = NewPrivateKeyFromString(seed)
		addressStrn = "ba5b87d2fb35fd1da29e91a851281c0a30f5a213"
	)
	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStrn, address.String())
	fmt.Println(address)
}
func TestPublicKeyToAddress(t *testing.T) {
	priKey := GeneratePrivateKey()
	pubKey := priKey.Public()
	address := pubKey.Address()
	assert.Equal(t, addressLen, len(address.Byte()))
	fmt.Println(address)
}
