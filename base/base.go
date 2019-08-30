package base

import (
	"bard/bard-plugin/util/CFB"
	"bard/bard-plugin/util/get"
	"crypto/des"
	"fmt"
	"log"
)
// 只开启了加密 AntiSniffing函数
const DefaultPri = 0x2111

const DefaultKey = "12345678"

const (
	ID = "base"
	Ver = "0.2.0"
	END_FLAG = "\r\n\r\n"
)

// 包初始化
func init() {
	// ParseConfig保证了一定有正确Config输出
	config, err := ParseConfig("./base_plugin_config.yml")
	if err != nil {
		log.Println(err)
	}
	BConfig = config
	V = Plugin{
		ID:  ID,
		Ver: Ver,
		Pri: BConfig.Priority,
		DESKEY: []byte(BConfig.DESKEY)[:des.BlockSize],
		END_FLAG: []byte(END_FLAG),
	}
}
var BConfig *Config

var V Plugin

type Plugin struct {
	ID  string
	Ver string
	Pri uint16				// pOAC
	DESKEY []byte
	END_FLAG []byte
}

func (p Plugin) EndCam() []byte {
	return p.END_FLAG
}

// send 为true时可以从返回值返回字节数组
// send 为false时只能从实参中返回字节数组

// 接收时， 返回byte数字是无效字段，仅int有效
func (p Plugin) Camouflage(bs []byte, send bool) ([]byte, int) {
	//fmt.Printf("%s: Camouflage\t %x\n", p.ID, p.Pri)
	if send {
		return get.Request(bs)
	} else {
		//    "Keep-Alive\r\n\r\n"
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
		return CFB.CFBEncrypter(cipherBlock, des.BlockSize, bs)
	} else {
		return CFB.CFBDecrypter(cipherBlock, des.BlockSize, bs)
	}
}

func (p Plugin) Ornament(bs []byte, send bool) ([]byte, int) {
	//fmt.Printf("%s: Ornament\t %x\n", p.ID, p.Pri)
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
