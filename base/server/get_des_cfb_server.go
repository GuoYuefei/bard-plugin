package main

import (
	"bard/bard-plugin/base"
	"bard/bard-plugin/util/get"
	"crypto/des"
)

var V = PluginServer{struct {
	ID     string
	Ver    string
	Pri    uint16
	DESKEY []byte
	END_FLAG []byte
	// 关闭了C函数
}{ID: string(base.ID), Ver: string(base.Ver), Pri: base.BConfig.Priority,
	DESKEY: []byte(base.BConfig.DESKEY)[:des.BlockSize], END_FLAG: []byte(base.END_FLAG)}}

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


