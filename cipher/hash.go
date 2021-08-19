package cipher

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
)

type HashCipher struct {
	src  []byte
	salt []byte
	algo hash.Hash
}
type HashCipherOpt func(*HashCipher)

func HashSalt(salt []byte) HashCipherOpt {
	return func(hc *HashCipher) {
		hc.salt = salt
	}
}

func HashAlgorithm(algo hash.Hash) HashCipherOpt {
	return func(hc *HashCipher) {
		hc.algo = algo
	}
}

func Hash(src []byte, opts ...HashCipherOpt) *HashCipher {
	hc := &HashCipher{
		src: src,
		// salt: []byte{0x19, 0x80, 0x03, 0x06, 0x20, 0x08, 0x03, 0x25},
		salt: []byte{},
		algo: md5.New(),
	}
	for _, opt := range opts {
		opt(hc)
	}
	return hc
}

func (hc *HashCipher) Bytes() ([]byte, error) {
	if _, err := hc.algo.Write(hc.src); err != nil {
		return nil, err
	}
	return hc.algo.Sum(hc.salt), nil
}

func (hc *HashCipher) String() (string, error) {
	bs, err := hc.Bytes()
	if err != nil {
		return "", fmt.Errorf("string: %w", err)
	}
	return string(bs), nil
}

func (hc *HashCipher) Hex() (string, error) {
	bs, err := hc.Bytes()
	if err != nil {
		return "", fmt.Errorf("hex: %w", err)
	}
	return hex.EncodeToString(bs), nil
}

func (hc *HashCipher) Base64() (string, error) {
	bs, err := hc.Bytes()
	if err != nil {
		return "", fmt.Errorf("base64: %w", err)
	}
	return base64.StdEncoding.EncodeToString(bs), nil
}
