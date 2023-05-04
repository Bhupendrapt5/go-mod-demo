package main

import (
	"fmt"

	"github.com/Bhupendrapt5/go-mod-demo/greetings"
)

func main() {
	message := greetings.Greetings("Gladys")
	fmt.Println(message)
}
