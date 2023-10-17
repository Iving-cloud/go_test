package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-protobuf/pb"
	"log"
)

func main() {
	p := &pb.Person{
		Name:   "timmy",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}

	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("matshal data len: %dn", len(data))

	np := &pb.Person{}
	if err = proto.Unmarshal(data, np); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("unmatshal person name: %s\n", np.Name)
	fmt.Println(np)
	fmt.Println(p.GetName())
	fmt.Println()
}
