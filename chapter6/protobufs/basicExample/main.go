package main

import (
	"fmt"

	pb "github.com/Hands-On-Restful-Web-services-with-Go/chapter6/protobufs/protofiles"
	"github.com/golang/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "Roger F",
		Email: "rf@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshalled proto data: ", body)
	fmt.Println("Unmarshalled struct: ", p1)
}
