package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	if isFormal {
		do(message)
	}
	do(alternate)
}

func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "hey!" + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}
func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
func PrintCustom(s string, custom string) {
	fmt.Println(s + custom)
}

func main() {
	var s = Salutation{}
	s.name = "Bob"
	s.greeting = "Hello"
	Greet(s, CreatePrintFunction("!!!"), false)
}
