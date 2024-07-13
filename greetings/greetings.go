package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for a named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name cannot be empty")
	}

	message := fmt.Sprintf("Hi %v, what up dawggggg!", name)
	return message, nil
}
