package main

import (
	"bytes"
	"fmt"
)

func main() {
	var bs = []byte("12430-8015623086789921434")
	lens := funindexlast(bs)
	fmt.Printf("外面： %s\n", bs[0:lens])
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