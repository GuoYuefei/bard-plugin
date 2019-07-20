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
}{ID: string("base"), Ver: string("0.1.0"), Pri: uint16(0x7111), DESKEY: []byte("12345678")}}

type PluginServer struct {
	base.Plugin
}
func (p PluginServer) Camouflage(bs []byte, send bool) ([]byte, int) {
	//fmt.Printf("%s: Camouflage\t %x\n", p.ID, p.Pri)
	if send {
		return get.Response(bs)
	} else {
		return get.Clear(bs)
	}
}


