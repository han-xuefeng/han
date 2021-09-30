package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main()  {

	test := &Student{
		Name: "han",
		Male: true,
		Scores: []int32{98,85,88},
	}

	data, err := proto.Marshal(test)

	if err != nil {
		log.Fatal("error :", err)
	}
	fmt.Println(data)
	fmt.Printf("%v.....%T", data,data)

	newTest := &Student{}

	proto.Unmarshal(data, newTest)
	fmt.Println(newTest.String())
}
