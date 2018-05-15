package main

import "./greeting"

func main() {

	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
	}
	salutations.Greet(greeting.CreatePrintFunction("?"), true, 6)
	// greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 6)
}
