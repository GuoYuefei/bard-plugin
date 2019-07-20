package get

import "bytes"

func Request(bs []byte) ([]byte, int){
	// 加头信息
	head := []byte(
		"GET /GuoYuefei/myjson/master/logo.png HTTP/1.1\r\n"+
			"Host: raw.githubusercontent.com\r\n"+
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36\r\n"+
			"Accept-Encoding: gzip, deflate, br\r\n"+
			"Connection: Keep-Alive\r\n\r\n")
	bs = append(head, bs...)
	return bs, len(bs)
}

func Response(bs []byte) ([]byte, int) {
	//
	head := []byte(
		"HTTP/1.1 200 OK\r\n"+
		"Transfer-Encoding: chunked\r\n"+
		"Server: GitHub.com\r\n"+
		"Status: 200 OK\r\n"+
		"Content-Encoding: gzip\r\n"+
		"Vary: Accept-Encoding\r\n\r\n")
	bs = append(head, bs...)
	return bs, len(bs)
}

func Clear(bs []byte) ([]byte, int) {
	var flag = []byte("\r\n\r\n")
	index := bytes.Index(bs, flag) + len(flag)
	if index < len(flag) {
		// 没找到
		return bs, len(bs)
	}
	var temp []byte = bs[index:]
	for k, v := range temp {
		bs[k] = v
	}
	return bs, len(temp)
}


