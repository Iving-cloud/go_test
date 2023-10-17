package main

//井字棋服务端
import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-protobuf/pb"
	"log"
	"net"
	"strings"
)

type Game_data struct {
	Used   uint8
	Client string
}

type pos_data struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Data struct {
	Client  string   `json:"client"`
	PosData pos_data `json:"pos"`
}

var ClientList []net.Conn
var Chess [3][3]Game_data

func main_game_handle(x int, y int, client string) int {
	//判断所在行
	if x > 2 || x < 0 || y > 2 || y < 0 {
		return -1
	}
	for i := 0; i < 2; i++ {
		if Chess[i][y].Used == 0 && Chess[i][y].Client != client {
			break
		}
	}
	if i == 2 {
		return 1
	}
	for i := 0; i < 2; i++ {
		if Chess[x][i].Used == 0 && Chess[x][i].Client != client {
			break
		}
	}
	if i == 2 {
		return 1
	}
	if (x+y)/2 != 0 {
		return 0
	}
	if Chess[1][1].Used == 0 && Chess[1][1].Client != client {
		return 0
	}
	if Chess[y][x].Used == 0 && Chess[y][x].Client != client {
		return 0
	}
	return 1
}

func handle_rev_msg(rev_msg string) {
	var d Data
	rev_msg = strings.ReplaceAll(rev_msg, "\x00", "")
	err := json.Unmarshal([]byte(rev_msg), &d)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return
	}
	// 输出解析后的结果
	fmt.Println("Client:", d.Client)
	//fmt.Println("Pos:", d.Pos)
	//handle_pos(d.Pos)
	fmt.Println("data:", d)
	Chess[d.PosData.X][d.PosData.Y].Used = 1
	Chess[d.PosData.X][d.PosData.Y].Client = d.Client
}

func handleConnect(conn net.Conn) {
	defer conn.Close()
	var con_tmp []net.Conn
	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Read error")
			for _, v := range ClientList {
				if v != conn {
					con_tmp = append(con_tmp, v)
				}
			}
			ClientList = con_tmp[:]
			fmt.Println("list:", conn)
			return
		}
		client := conn.RemoteAddr()
		fmt.Println("client", client)
		fmt.Println("Receive data", string(buffer))

		for _, v := range ClientList {
			if v != conn {
				v.Write([]byte(string(buffer)))
			}
		}
		//conn.Write(send_msg())
		str := string(buffer)
		handle_rev_msg(str)
	}
}

func send_msg() []byte {
	p := &pb.Person{
		Name:   "timmy",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DATA", data)
	return []byte(data)
}

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	//str := `{"client":"a","pos":"1,2"}`
	//handle_rev_msg(str)
	defer listener.Close()
	fmt.Println("Server listening on localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		ClientList = append(ClientList, conn)
		go handleConnect(conn)
	}
}
