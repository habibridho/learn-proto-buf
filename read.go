package main

import (
	"io/ioutil"
	"log"
	"learn-proto-buf/model"
	"github.com/golang/protobuf/proto"
	"fmt"
)

func main() {
	in, err := ioutil.ReadFile("test")

	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	book := &model.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Printf("Yow %v\n", book)
}
