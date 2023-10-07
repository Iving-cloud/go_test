package main

import (
	"fmt"
	"net"
)

func handleConnect(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error")
		return
	}
	fmt.Println("Receive data",string(buffer))
	conn.Write([]byte("Receive data."))
}

func main() {
	listener, err := net.Listen("tcp",":8888")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on localhost:8080")
	for{
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		go handleConnect(conn)
	}
}