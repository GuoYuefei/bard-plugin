package util

import (
	"net"
	"runtime"
)

// 出错 or 读满bs结束
func ReadFull(conn net.Conn, bs []byte) (n int, err error) {
	lens := len(bs)
	n = 0
	for n != lens {
		i, err := conn.Read(bs[n:])
		n += i
		if err != nil {
			return n, err
		}
		if n == lens {
			return n, nil
		}
		// 没读取完就让出时间片等下在读
		runtime.Gosched()
	}
	return lens, nil
}
