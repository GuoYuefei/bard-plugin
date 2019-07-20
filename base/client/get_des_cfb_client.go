package main

import "bard-plugin/base"

var V = base.Plugin{
	ID:  "base",
	Ver: "0.1.0",
	Pri: 0x7111,
	DESKEY: []byte("12345678"),
}

//func main() {
//	ws := []byte("cctv.com12345678910")
//	ws, l :=V.AntiSniffing(ws, true)
//	fmt.Printf("%d \t  %s\n", l, ws)
//	_, l = V.AntiSniffing(ws, false)
//	fmt.Printf("%d \t %s\n", l, ws[:l])
//
//}


