package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// No need to specify the variable name since it's not going to be used
func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi there!"
}

// No need to specify the variable name since it's not going to be used
func (spanishBot) getGreeting() string {
	// Very custom logic for generating a spanish greeting
	return "Hola!"
}
