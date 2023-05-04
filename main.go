package main

import (
	"fmt"
	"log"

	"github.com/Bhupendrapt5/go-mod-demo/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Greetings(names)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(messages)
}
