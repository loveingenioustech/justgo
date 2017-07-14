package main

import (
	"errors"
	"net/rpc"
	"net/http"
	"fmt"
	"net"
	"os"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	//demoHttpServer()

	//demoTcpServer()

	demoJsonServer()
}

func demoHttpServer(){
	arith := new(Arith)
	rpc.Register(arith) // 注册一个Arith的RPC服务
	rpc.HandleHTTP() // 把该服务注册到了HTTP协议上

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func demoTcpServer(){
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkServerError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkServerError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func checkServerError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}


func demoJsonServer(){
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkServerError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkServerError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}