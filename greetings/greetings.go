package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for a named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name cannot be empty")
	}

	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomGreeting() string {
	greetings := []string{
		"Hi %v, what up!",
		"Hello %v",
		"YOOO %v!",
	}

	return greetings[rand.Intn(len(greetings))]
}
