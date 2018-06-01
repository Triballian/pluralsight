package main

import "fmt"

func proceqtwo(e *eq, o *op) {
	var lo op
	var ro op

	e.delimeters = append(e.delimeters, "=")

	o.delimeters = append(o.delimeters, "+", "-")
	lo.delimeters = append(lo.delimeters, "x", "y")

	ro.delimeters = append(lo.delimeters, "x", "y")
	e.delSpace()
	split(e, e.expression)

	split(o, lo.expression)

	o.lsplit(e)
	o.rsplit(e)
	o.expression = e.expression
	o.expression = e.expression
	lo.expression = o.lSide
	ro.expression = o.rSide
	split(&lo, lo.expression)
	split(&ro, ro.expression)
	fmt.Printf("o.lside is: %v\n", o.lSide)
	fmt.Printf("lo.rside is: %v\n", ro.rSide)
	var h head
	if len(e.lSide) > 0 {
		h.LsideEquation = true

		if len(o.lSide) > 0 {
			h.lsideOperator = true
			o.termSplit(&h, e)
		}
	}
	// splitEq(o)
	// splitOp(o)
	// o.split("+")
	// 	e.terms[0].coeficient = append(e.terms[0].coeficient, e.equation)
	// 	fmt.Printf("eqution: %s", e.equation)
	// }

}

// func splitEq(ex expression) {
// 	ex.split("=")
// }
