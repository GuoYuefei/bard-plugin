package main

import (
	"bard-plugin/base"
	"fmt"
)

var V = base.Plugin{
	ID:  "base",
	Ver: "0.1.0",
	Pri: 0x3111,
	DESKEY: []byte("12345678"),
	END_FLAG: []byte("\r\n\r\n"),
}

func main() {
	var l int
	ws := []byte("h23t43Keep-Alive\r\n\r\n\r\n\r\n4tp5453s://cctv.com12345678910")
	fmt.Printf("%d \t  %s\n", len(ws), ws)

	//ws, l = V.AntiSniffing(ws, true)
	//fmt.Printf("a:\t%d \t  %s\n", l, ws)

	ws, l = V.Camouflage(ws, true)
	fmt.Printf("ac:\t%d \t  %s\n", l, ws)

	_, l = V.Camouflage(ws, false)
	fmt.Printf("a:\t%d \t  %s\n", l, ws[:l])

	//_, l = V.AntiSniffing(ws[:l], false)
	//fmt.Printf("%d \t %s\n", l, ws[:l])

}


