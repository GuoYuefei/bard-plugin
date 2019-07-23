package main

import (
	"bytes"
	"fmt"
)

func main() {
	var bs = []byte("12430-8015623086789921434")
	lens := funindexlast(bs)
	fmt.Printf("外面： %s\n", bs[0:lens])
	ll()
}

func fun(bs []byte) int {
	temp := bytes.TrimLeft(bs, "12430-")
	for k, v := range temp {
		bs[k] = v
	}

	fmt.Printf("里面： %s\n", bs)
	return len(temp)
}

func funindexlast(bs []byte) int {
	index := bytes.LastIndex(bs, []byte("30-"))
	var temp []byte = bs[index+3:]
	for k, v := range temp {
		bs[k] = v
	}
	fmt.Printf("里面： %s\n", bs)
	return len(temp)
}

func ll() {
	var ll = []byte("1234567890")
	index := bytes.Index(ll[3:], []byte("90"))
	fmt.Println(index)
	index = bytes.Index(ll, []byte("90"))
	fmt.Println(index, string(ll[0:index]))
	fmt.Println(ll[10:])
}