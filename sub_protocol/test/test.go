package test

/**
	仅仅为了测试，请勿使用
	与主程序中default基本相同，多了一个打印
 */

import (
	"bard/bard-plugin/sub_protocol"
	"bard/bard-plugin/sub_protocol/util"
	"fmt"
	"net"
)

func readDo(conn net.Conn) ([]byte, int) {
	// default len is two byte
	lslice := make([]byte, 2)
	_, err := util.ReadFull(conn, lslice)
	if err != nil {
		return nil, 0
	}
	// 大端
	lenh, lenl := int(lslice[0]), int(lslice[1])
	l := lenh<<8+lenl
	fmt.Println("it's test!")
	return lslice, l
}

func writeDo(bs []byte) ([]byte, int) {
	l := len(bs)
	lenh, lenl := byte(l>>8), byte(l)
	//fmt.Println(lenh, lenl)
	lslice := []byte{lenh, lenl}
	bs = append(lslice, bs...)
	return bs, len(bs)
}


var T = sub_protocol.NewAssembleTCSP(
	"test",
	readDo,
	writeDo,
	)

