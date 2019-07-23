package get

import (
	"bytes"
	"fmt"
	"strconv"
)

func Request(bs []byte) ([]byte, int) {
	// 加头信息
	head := []byte(
		"POST /GuoYuefei/myjson/master/logo.png HTTP/2.0\r\n" +
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
		"HTTP/2.0 200 OK\r\n" +
			//"Transfer-Encoding: chunked\r\n" +
			"access-control-allow-methods: POST\r\n"+
			"access-control-max-age: 3600\r\n"+
			"Server: GitHub.com\r\n" +
			"Status: 200 OK\r\n" +
			"server: AmazonS3\r\n"+
			"accept-ranges: bytes\r\n"+
			"Content-Encoding: gzip\r\n" +
			"Content-Length: " + strconv.Itoa(len(bs)) + "\r\n" +
			"Vary: Accept-Encoding\r\n\r\n")
	bs = append(head, bs...)
	return bs, len(bs)
}


func Clear(bs []byte) ([]byte, int) {
	// bs是混淆的头部，需要解析后告诉主程序剩下的多少大小的数据才是真正的负载
	var cl = []byte("Content-Length: ")

	lengthBegin := bytes.Index(bs, cl) + len(cl)
	lengthEnd := bytes.Index(bs[lengthBegin:], []byte("\r\n")) + lengthBegin
	length, e := strconv.Atoi(string(bs[lengthBegin:lengthEnd]))
	if e != nil {
		fmt.Println("报文格式不对")
	}

	return bs, length

}
