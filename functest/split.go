package main

import "fmt"

func (e *eq) split() {
	for i, c := range e.expression {
		cc := fmt.Sprintf("%c", c)
		if cc == "=" {
			e.delimeter = cc
			e.lSide = e.expression[:i]
			e.rSide = e.expression[i+1:]

		}
	}
}
func (o *op) lsplit(e *eq) {
	for i, c := range e.expression {
		cc := fmt.Sprintf("%c", c)
		if (cc == "+") || (cc == "-") {
			o.loperator = true
			o.delimeter = cc
			o.lSide = e.expression[:i]
			o.rSide = e.expression[i+1:]

		}
	}
}

func (o *op) rsplit(e *eq) {
	for i, c := range e.expression {
		cc := fmt.Sprintf("%c", c)
		if (cc == "+") || (cc == "-") {
			o.delimeter = cc
			o.lSide = e.expression[:i]
			o.rSide = e.expression[i+1:]

		}
	}
}
