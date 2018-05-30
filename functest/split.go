package main

import "fmt"

//unc (e *eq) split() {
func split(d delimeter) {
	for i, c := range d.getexp() {
		cc := fmt.Sprintf("%c", c)
		if d.isdelimeter(cc) {
			d.setside(cc, i)
			// d.delimeter = cc
			// d.lSide = e.expression[:i]
			// d.rSide = e.expression[i+1:]

		}
	}
}
func (e *eq) getexp() (exp string) {
	exp = e.expression
	return
}
func (o *op) getexp() (exp string) {
	exp = o.expression
	return
}

func (e *eq) setside(cc string, i int) {
	e.delimeter = cc
	e.lSide = e.expression[:i]
	e.rSide = e.expression[i+1:]
}

func (o *op) setside(cc string, i int) {
	o.delimeter = cc
	o.lSide = o.expression[:i]
	o.rSide = o.expression[i+1:]
}

// do this with expression as well
func (o *eq) isdelimeter(cc string) (b bool) {
	if cc == "=" {
		b = true
		return
	}
	return
}
func (e *op) isdelimeter(cc string) (b bool) {
	if (cc == "+") || (cc == "-") {
		b = true
		return
	}
	return
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
