package main

import "fmt"

var test = []struct {
	baby string
	dog  string
}{
	{"baby", "blue"},
	{"dog", "ray"},
}

func main() {
	fmt.Printf("this is the output %v", test)
}
