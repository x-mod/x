package cipher

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashCipher_Base64(t *testing.T) {
	s1, err := Hash([]byte("thisistest")).Base64()
	assert.Nil(t, err)
	log.Println("s1:", s1)

	s2, err := Hash([]byte("thisistest")).Base64()
	assert.Nil(t, err)
	log.Println("s2:", s2)

	assert.Equal(t, s1, s2)
	s3, err := Hash([]byte("thisistest1")).Base64()
	assert.Nil(t, err)
	log.Println("s3:", s3)
	assert.NotEqual(t, s1, s3)

	cipher, err := NewAES256CFBCipher(Secret([]byte("adjflakdfja")), Salt([]byte("123456")))
	assert.Nil(t, err)

	enc, err := cipher.Encrypt([]byte("this is original data "))
	assert.Nil(t, err)
	log.Println("encode:", string(enc))
	henc := hex.EncodeToString(enc)
	log.Println("hex encode:", henc)

	hdec, err := hex.DecodeString(henc)
	assert.Nil(t, err)
	log.Println("hex decode:", hdec)

	dec, err := cipher.Decrypt(enc)
	assert.Nil(t, err)
	log.Println("decode:", string(dec))

	ns, err := cipher.Encrypt([]byte("djsaklfjaldfjalskjdd---3424242fdkasldfjasdfjasdfaewafsafdeasdfsf"))
	assert.Nil(t, err)
	log.Println("encode num:", string(ns))
	log.Println("encode num hex:", hex.EncodeToString(ns))

	ne, err := cipher.Decrypt(ns)
	assert.Nil(t, err)
	assert.Equal(t, string(ne), "djsaklfjaldfjalskjdd---3424242fdkasldfjasdfjasdfaewafsafdeasdfsf")
	// log.Println("decode num:", string(ne))
}
