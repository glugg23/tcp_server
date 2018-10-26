package main

import (
	"fmt"
	"net"
)

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Println(string(buf[:reqLen]))
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("Listening Error")
		return
	}

	defer l.Close() //Close listener when program quits

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Accepting Error")
			return
		}

		go handleRequest(conn)
	}
}
