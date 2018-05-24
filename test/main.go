package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	s = "2*x**2+3y**2"
	seq := strings.Split(s, `=`)
	fmt.Println("seq[0] has the value of %s", seq[1])
}
