package main

import (
	"bard-plugin/base"
	"bard-plugin/util/get"
)

var V = PluginServer{struct {
	ID     string
	Ver    string
	Pri    uint16
	DESKEY []byte
	END_FLAG []byte
}{ID: string("base"), Ver: string("0.1.0"), Pri: uint16(0x5111), DESKEY: []byte("12345678"), END_FLAG: []byte("\r\n\r\n")}}

type PluginServer struct {
	base.Plugin
}
func (p PluginServer) Camouflage(bs []byte, send bool) ([]byte, int) {
	//fmt.Printf("%s: Camouflage\t %x\n", p.ID, p.Pri)
	if send {
		return get.Response(bs)
	} else {
		return get.Clear(bs, []byte("\r\n\r\n"))
	}
}


