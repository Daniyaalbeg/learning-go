package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(3)

	// message, err := greetings.Hello("Dan")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	messages, err := greetings.Hellos([]string{
		"Daniyaal",
		"Kaamil",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
