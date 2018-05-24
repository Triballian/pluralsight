package main

import "fmt"

type eq struct {
	header    *head
	delimeter string
	leftSide  string
	rightSide string
	opLeft    string
	opRight   string
	terms     []*term
	operators []*op

	// coefficient terms or constant terms

}

type expression interface {
	split(delimeter string)
}

type head struct {
	expression string
}
type op struct {
	header    *head
	delimeter string
	lside     string
	rside     string
	operator  string
}
type term struct {
	lside      bool
	operator   string   // left side, right side of operator or there is no operator adjecent to this term.
	coeficient []string // every term has a coeffecient even if the coefficient is 1
	variable   rune     // x, y, or nil
	constant   int
	exponent   int // even if the exponent is zero
}

func main() {
	var h head
	var e eq
	var o op
	h.expression = prmptforInput()
	e.header = &h
	o.header = &h
	// fmt.Printf("equation : %s", e.value)
	// seq := strings.Split(pe.equation, `=`)
	proceqtwo(&e, &o)
	fmt.Printf("equation : %s", h.expression)

	// fmt.Printf("Left side is %s\n", seq[0])
	fmt.Printf("left side: %s", e.leftSide)
	fmt.Printf("right side: %s", e.rightSide)
	fmt.Printf("delimeter: %s", e.delimeter)
	fmt.Printf("op leftside: %s", o.lside)

}
