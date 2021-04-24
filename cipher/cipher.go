package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

type AES256CFBCipher struct {
	secret []byte
	block  cipher.Block
	enc    cipher.Stream
	dec    cipher.Stream
}
type AES256CFBCipherOpt func(*AES256CFBCipher)

func Secret(secret []byte) AES256CFBCipherOpt {
	return func(ac *AES256CFBCipher) {
		ac.secret = secret
	}
}
func HexString(secret string) AES256CFBCipherOpt {
	return func(ac *AES256CFBCipher) {
		sec, err := hex.DecodeString(secret)
		if err != nil {
			panic(err)
		}
		ac.secret = sec
	}
}

func NewAES256CFBCipher(opts ...AES256CFBCipherOpt) (*AES256CFBCipher, error) {
	c := &AES256CFBCipher{}
	for _, opt := range opts {
		opt(c)
	}
	key := secretToKey(c.secret, 32)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := key[:aes.BlockSize]
	c.enc = cipher.NewCFBEncrypter(block, iv)
	c.dec = cipher.NewCFBDecrypter(block, iv)
	c.block = block
	return c, nil
}

func secretToKey(secret []byte, size int) []byte {
	// size mod 16 must be 0
	h := md5.New()
	buf := make([]byte, size)
	count := size / md5.Size
	// repeatly fill the key with the secret
	for i := 0; i < count; i++ {
		h.Write(secret)
		copy(buf[md5.Size*i:md5.Size*(i+1)-1], h.Sum(nil))
	}
	return buf
}

func (c *AES256CFBCipher) Encrypt(origData []byte) ([]byte, error) {
	encrypted := make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(c.block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted, nil
}

func (c *AES256CFBCipher) Decrypt(encrypted []byte) ([]byte, error) {
	if len(encrypted) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(c.block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted, nil
}

func (c *AES256CFBCipher) Encrypter() cipher.Stream {
	return c.enc
}
func (c *AES256CFBCipher) Decrypter() cipher.Stream {
	return c.dec
}
func (c *AES256CFBCipher) DecryptReader(rd io.Reader) *cipher.StreamReader {
	return &cipher.StreamReader{
		S: c.dec,
		R: rd,
	}
}
func (c *AES256CFBCipher) EncryptWriter(wr io.Writer) *cipher.StreamWriter {
	return &cipher.StreamWriter{
		S: c.enc,
		W: wr,
	}
}
