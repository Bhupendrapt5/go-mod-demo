package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Greeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("gimme a name dumbass")
	}

	greetMSg := fmt.Sprintf(randomFormat(), name)

	return greetMSg, nil
}

func Greetings(names []string) (map[string]string, error) {
	greetingMsgs := make(map[string]string)

	for _, name := range names {

		message, err := Greeting(name)
		if err != nil {
			return nil, err
		}

		greetingMsgs[name] = message

	}

	return greetingMsgs, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {

	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
