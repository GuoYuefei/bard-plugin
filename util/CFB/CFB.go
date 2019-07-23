package CFB

import (
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// send = true
func CFBEncrypter(block cipher.Block, blockSize int, bs []byte) ([]byte, int) {
	// 加密
	ciphertext := make([]byte, blockSize+len(bs))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		//panic(err)
		return nil, 0
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[blockSize:], bs)
	return ciphertext, len(ciphertext)
}

// send = false
func CFBDecrypter(block cipher.Block, blockSize int, bs []byte) ([]byte, int) {
	// 解密
	if len(bs) < blockSize {
		//panic("ciphertext too short")
		return nil, 0
	}
	iv := make([]byte, 8)
	l := len(bs) - blockSize
	copy(iv, bs[:blockSize])
	copy(bs[:l], bs[blockSize:])
	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(bs, bs)
	return bs, l
}
