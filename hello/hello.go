package main

import "./greeting"

func main() {
	// var s = greeting.Salutation{}
	// s.Name = "Bob"
	// s.Greeting = "Hello"

	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
	}
	slice = slice[1:2]
	greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 6)
}
