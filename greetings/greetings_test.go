package greetings

import (
	"regexp"
	"testing"
)

// Test Greeting with valid name
func TestGreetingName(t *testing.T) {
	name := "Morpheus"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Greeting(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Morpheus") = %q, %v want match for %q, nil`, msg, err, want)
	}

}

// Test Greeting with empty name
func TestGreetingEmptyName(t *testing.T) {

	msg, err := Greeting("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v want "", nil`, msg, err)
	}

}
