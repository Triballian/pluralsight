package main

func proceqtwo(e *eq, o *op) {

	e.delSpace()
	splitEq(e)
	splitOp(o)
	// o.split("+")
	// 	e.terms[0].coeficient = append(e.terms[0].coeficient, e.equation)
	// 	fmt.Printf("eqution: %s", e.equation)
	// }

}
func (o *op) split(delimeter string) {}
func splitEq(ex expression) {
	ex.split("=")
}
func splitOp(ex expression) {
	ex.split("+")
}
