package main

import (
	"fmt"
	"log"
	"github.com/sarthakvk/greet"
)

func main() {
	//setting logger
	//	log.SetPrefix("greetings: ")
	//	log.SetFlags(0)

	message, err := greet.Hello("Sartak")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
