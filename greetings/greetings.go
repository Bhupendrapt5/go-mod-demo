package greetings

import "fmt"

func Greetings(name string) string {

	greetMSg := fmt.Sprintf("hello %v. How are you doing", name)

	return greetMSg
}
