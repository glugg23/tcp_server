package main

import (
	"fmt"
	"net"
)

func handleRequest(conn net.Conn) {
	fmt.Println("Opened connected")

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			//Need to make sure this only happens when there are no more packets
			fmt.Println("Closed connection")
			break
		}

		fmt.Println(string(buf[:n]))
	}

	conn.Close()
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("Listening Error")
		return
	}

	fmt.Println("Started listening")

	defer l.Close() //Close listener when program quits

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Accept Error")
			continue
		}

		go handleRequest(conn)
	}

	fmt.Println("Finished listening")
}
