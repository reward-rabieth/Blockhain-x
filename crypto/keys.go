package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io"
)

const (
	signatureLen = 64
	privKeyLen   = 64
	pubKeyLen    = 32
	seedLen      = 32
	addressLen   = 20
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func NewPrivateKeyFromString(s string) *PrivateKey {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return NewPrivateKeyFromSeed(b)
}
func NewPrivateKeyFromSeed(seed []byte) *PrivateKey {
	if len(seed) != seedLen {
		panic("invalid seed length, must be 32")
	}
	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedLen)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}
	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.key, msg),
	}
}
func (p *PrivateKey) Public() *PublicKey {
	b := make([]byte, pubKeyLen)
	copy(b, p.key[32:])
	return &PublicKey{
		Key: b,
	}
}

type PublicKey struct {
	Key ed25519.PublicKey
}

func PublicKeyFromBytes(b []byte) *PublicKey {
	if len(b) != pubKeyLen {
		panic("invalid public key length")
	}
	return &PublicKey{
		Key: ed25519.PublicKey(b),
	}

}
func (p *PublicKey) Address() Address {
	return Address{
		value: p.Key[len(p.Key)-addressLen:],
	}

}

func (p *PublicKey) Bytes() []byte {
	return p.Key
}

type Signature struct {
	value []byte
}

func SignatureFromBytes(b []byte) *Signature {
	if len(b) != signatureLen {
		panic("length of bytes not equal to 64")
	}
	return &Signature{
		value: b,
	}
}

func (s *Signature) Bytes() []byte {
	return s.value
}

func (s *Signature) Verify(pubKey *PublicKey, msg []byte) bool {
	return ed25519.Verify(pubKey.Key, msg, s.value)
}

type Address struct {
	value []byte
}

func (a Address) Byte() []byte {
	return a.value
}

func (a Address) String() string {
	return hex.EncodeToString(a.value)
}
