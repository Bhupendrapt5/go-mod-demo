package main

import (
	"fmt"
	"log"

	"github.com/Bhupendrapt5/go-mod-demo/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Greetings("")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(message)
}
