package main

import "fmt"

type ex struct {
	equation eq
	operator []op
}
type eq struct {
	expression string
	delimeter  string
	leftSide   string
	rightSide  string
	opLeft     string
	opRight    string
	terms      []*term
	operators  []*op

	// coefficient terms or constant terms

}

type expression interface {
	split(delimeter string)
}

type lop struct {
}
type rop struct {
}

type op struct {
	delimeter string
	lterm     string
	rterm     string
	operator  string
	lside     bool
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
	var e eq
	var o op

	e.expression = prmptforInput()
	// fmt.Printf("equation : %s", e.value)
	// seq := strings.Split(pe.equation, `=`)
	proceqtwo(&e, &o)
	fmt.Printf("equation : %s", e.expression)

	// fmt.Printf("Left side is %s\n", seq[0])
	fmt.Printf("left side: %s", e.leftSide)
	fmt.Printf("right side: %s", e.rightSide)
	fmt.Printf("delimeter: %s", e.delimeter)

}
