package main

import "./greeting"

func main() {
	var s = greeting.Salutation{}
	s.Name = "Bob"
	s.Greeting = "Hello"
	greeting.Greet(s, greeting.CreatePrintFunction("!!!"), true, 6)
}
