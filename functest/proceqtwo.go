package main

func proceqtwo(e *eq, o *op) {

	e.delSpace()
	e.split()
	o.lsplit(e)
	o.rsplit(e)
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
