package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

const (
	Language = "Go"
	A        = iota
	B        = iota
	C        = iota
)

func main() {
	var s = Salutation{}
	s.name = "Bob"
	fmt.Println(A, B, C)
	fmt.Println(Language)
}
