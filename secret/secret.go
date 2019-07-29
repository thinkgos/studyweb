// package secret imp AES-128 CBC PCKS5 Padding with your salt

package secret

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"errors"
)

var errInputNotFullBlocks = errors.New("input not full blocks")
var errUnPaddingOutOfRange = errors.New("UnPadding out of range")

type Secret struct {
	salt string
}

func NewSpecial(salt string) *Secret {
	return &Secret{salt}
}

func pcks5Padding(origData []byte, size int) []byte {
	padSize := size - len(origData)%size
	padText := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(origData, padText...)
}

func pcks5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unPadSize := int(origData[length-1])
	if unPadSize > length {
		return nil, errUnPaddingOutOfRange
	}
	return origData[:(length - unPadSize)], nil
}

func (this *Secret) Encrypt(key, origData []byte) []byte {
	newKey := md5.Sum(key)
	iv := md5.Sum(append([]byte(this.salt), key...))

	block, err := aes.NewCipher(newKey[:])
	if err != nil {
		panic(err) // never happen
	}
	orig := pcks5Padding(origData, block.BlockSize())
	out := make([]byte, len(orig))
	cipher.NewCBCEncrypter(block, iv[:]).CryptBlocks(out, orig)
	return out
}
func (this *Secret) Decrypt(key, origData []byte) ([]byte, error) {
	newKey := md5.Sum(key)
	iv := md5.Sum(append([]byte(this.salt), key...))

	block, err := aes.NewCipher(newKey[:])
	if err != nil {
		panic(err) // never happen
	}
	if len(origData)%block.BlockSize() != 0 {
		return nil, errInputNotFullBlocks
	}
	out := make([]byte, len(origData))
	cipher.NewCBCDecrypter(block, iv[:]).CryptBlocks(out, origData)
	return pcks5UnPadding(out)
}
