package main

import (
	"fmt"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("无法连接到服务器:", err)
		return
	}
	defer conn.Close()
	message :="hello sever!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("发送数据失败", err)
		return
	}
	fmt.Println("成功发送数据到服务器:", message)
}