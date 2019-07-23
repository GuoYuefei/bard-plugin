package get

import (
	"bytes"
	"fmt"
	"strconv"
)

func Request(bs []byte) ([]byte, int) {
	// 加头信息
	head := []byte(
		"POST /GuoYuefei/myjson/master/logo.png HTTP/1.1\r\n" +
			"Host: raw.githubusercontent.com\r\n" +
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36\r\n" +
			"Accept-Encoding: gzip, deflate, br\r\n" +
			"Content-Length: " + strconv.Itoa(len(bs)) + "\r\n" +
			"Content-Type: image/jpeg\r\n" +
			"Connection: Keep-Alive\r\n\r\n")
	bs = append(head, bs...)
	return bs, len(bs)
}

func Response(bs []byte) ([]byte, int) {
	//
	head := []byte(
		"HTTP/1.1 200 OK\r\n" +
			"Transfer-Encoding: chunked\r\n" +
			"Server: GitHub.com\r\n" +
			"Status: 200 OK\r\n" +
			"Content-Encoding: gzip\r\n" +
			"Content-Length: " + strconv.Itoa(len(bs)) + "\r\n" +
			"Vary: Accept-Encoding\r\n\r\n")
	bs = append(head, bs...)
	return bs, len(bs)
}

func Clear(bs []byte, Splitter []byte) ([]byte, int) {
	// 二维切片记录分块内容的  指针方式 100应该不会再多了
	var temp [][]byte = make([][]byte, 100)
	var blocksize []byte = make([]byte, 200) // node 每块的长度， 每一块长度占两字节  大端存取
	var cl = []byte("Content-Length: ")
	var flag = Splitter
	var t = bs
	var i byte = 0
	for len(t) != 0 { // break条件
		var lengthBegin int
		var lengthEnd int
		var length int
		var e error
		index := bytes.Index(t, flag) + len(flag)
		if index < len(flag) {
			// 没找到
			//return t, len(t)
			goto BREAK
		}
		lengthBegin = bytes.Index(t, cl) + len(cl)
		lengthEnd = bytes.Index(t[lengthBegin:], []byte("\r\n")) + lengthBegin
		length, e = strconv.Atoi(string(t[lengthBegin:lengthEnd]))

		if e != nil {
			fmt.Println("插件格式错误:", e)
		}

		// 可能剩余的数据块没达到长度的
		if length > len(t[index:]) {
			goto BREAK
		}
		temp[i] = t[index : index+length]
		blocksize[2*i] = byte(len(temp[i]) / 256)
		blocksize[2*i+1] = byte(len(temp[i]) % 256)
		t = t[index+length:]
		i++
		continue
	BREAK:
		blocksize[2*i] = byte(len(t) / 256)
		blocksize[2*i+1] = byte(len(t) % 256)
		temp[i] = t
		i++
		break
	}
	var index int = 0 //最后为长度
	for j := byte(0); j < i; j++ {

		for _, v := range temp[j] {

			bs[index] = v
			index++
		}
	}
	blocksize = append([]byte{i}, blocksize...)

	// 返回各块的大小但第一个自己表示分了多少块和处理完数据的长度+未处理数据的总长 也就是blocksize中各数据（除第一字节）相加=index !没处理完的数据是因为数据不够，还需要从网络流中读取
	return blocksize, index
}
