package main

import (
	"learn-proto-buf/model"
	"github.com/golang/protobuf/proto"
	"log"
	"io/ioutil"
)

func main() {
	book := &model.AddressBook{}
	person := &model.Person{}
	person.Email = "habibridho@live.com"
	person.Name = "Habib"
	book.People = append(book.People, person)

	out, err := proto.Marshal(book)

	if err != nil {
		log.Fatalln("Failed encode address book:", err)
	}

	if err := ioutil.WriteFile("test", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
