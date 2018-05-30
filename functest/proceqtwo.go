package main

import "fmt"

func proceqtwo(e *eq, o *op) {

	e.delSpace()
	split(e)

	o.lsplit(e)
	o.rsplit(e)
	fmt.Printf("o.lside is %v", o.lSide)
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
