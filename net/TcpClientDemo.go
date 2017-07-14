package main

import (
	"net"
	"io/ioutil"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service) // 获得tcpAddr
	checkClientError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr) // 创建TCP连接conn
	checkClientError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkClientError(err)

	result, err := ioutil.ReadAll(conn)
	checkClientError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

func checkClientError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}