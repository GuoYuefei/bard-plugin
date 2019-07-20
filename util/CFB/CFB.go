package CFB

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"io"
)

// send = true
func CFBEncrypter(block cipher.Block, bs []byte) ([]byte, int) {
	// 加密
	ciphertext := make([]byte, des.BlockSize+len(bs))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		//panic(err)
		return nil, 0
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[des.BlockSize:], bs)
	return ciphertext, len(ciphertext)
}

// send = false
func CFBDecrypter(block cipher.Block, bs []byte) ([]byte, int) {
	// 解密
	if len(bs) < des.BlockSize {
		//panic("ciphertext too short")
		return nil, 0
	}
	iv := make([]byte, 8)
	l := len(bs) - des.BlockSize
	copy(iv, bs[:des.BlockSize])
	copy(bs[:l], bs[des.BlockSize:])
	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(bs, bs)
	return bs, l
}
