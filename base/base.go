package base

import (
	"bard-plugin/util/CFB"
	"bard-plugin/util/get"
	"crypto/des"
	"fmt"
)

type Plugin struct {
	ID  string
	Ver string
	Pri uint16
	DESKEY []byte
}

func (p Plugin) EndCam() byte {
	return 0x0ff
}

// send 为true时可以从返回值返回字节数组
// send 为false时只能从实参中返回字节数组

func (p Plugin) Camouflage(bs []byte, send bool) ([]byte, int) {
	//fmt.Printf("%s: Camouflage\t %x\n", p.ID, p.Pri)
	if send {
		return get.Request(bs)
	} else {
		return get.Clear(bs)
	}
}

func (p Plugin) AntiSniffing(bs []byte, send bool) ([]byte, int) {
	//fmt.Printf("%s: AntiSniffing\t %x\n", p.ID, p.Pri)
	cipherBlock, err := des.NewCipher(p.DESKEY)
	if err != nil {
		fmt.Printf("des error")
		return nil, 0
	}
	if send {
		// 加密
		return CFB.CFBEncrypter(cipherBlock, bs)
	} else {
		return CFB.CFBDecrypter(cipherBlock, bs)
	}
}

func (p Plugin) Ornament(bs []byte, send bool) ([]byte, int) {
	fmt.Printf("%s: Ornament\t %x\n", p.ID, p.Pri)
	return bs, len(bs)
}

func (p Plugin) Priority() uint16 {
	return p.Pri
}

func (p Plugin) GetID() string {
	return p.ID
}

func (p Plugin) Version() string {
	return p.Ver
}
