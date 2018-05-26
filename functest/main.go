package main

import "fmt"

// if the head is on the left side and there is an o.operator. Then run terSlide on o.lside
// delimeters are = , + || -, var, * || **
type eq struct {
	expression string
	// header    *head
	delimeter string
	lSide     string
	rSide     string
	opLeft    string
	opRight   string
	lterms    []*term
	rterms    []*term
	operators []*op

	// coefficient terms or constant terms

}

type head struct {
	LsideEquation bool
	lsideOperator bool
	lsideOVar     bool
}

type variable struct {
	lside string
	rside string
}

type op struct {
	delimeter string
	lSide     string
	rSide     string
	operator  string
	loperator bool
}
type term struct {
	lside      bool
	operator   string   // left side, right side of operator or there is no operator adjecent to this term.
	coeficient []string // every term has a coeffecient even if the coefficient is 1
	variable   string   // x, y, or nil
	constant   int
	baseEx     string
	varEx      string
	base       string
}

func main() {
	// var h head
	var e eq
	var o op
	input := prmptforInput()
	e.expression = input
	// fmt.Printf("equation : %s", e.value)
	// seq := strings.Split(pe.equation, `=`)
	proceqtwo(&e, &o)
	fmt.Printf("equation : %s", e.expression)

	// fmt.Printf("Left side is %s\n", seq[0])
	fmt.Printf("left side: %s", e.lSide)
	fmt.Printf("right side: %s", e.rSide)
	fmt.Printf("delimeter: %s", e.delimeter)
	fmt.Printf("op leftside: %s", o.lSide)

}
