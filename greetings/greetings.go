package greetings

import (
	"errors"
	"fmt"
)

func Greetings(name string) (string, error) {
	if name == "" {
		return "", errors.New("gimme a name dumbass")
	}

	greetMSg := fmt.Sprintf("hello %v. How are you doing", name)

	return greetMSg, nil
}
