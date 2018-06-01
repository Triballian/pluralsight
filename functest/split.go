package main

import "fmt"

func split(d delimeter, expression string) {
	for i, c := range expression {
		cc := fmt.Sprintf("%c", c)
		if d.isdelimeter(cc) {
			d.setside(cc, i)
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
	if len(o.lSide) > 0 {

	}
	o.delimeter = cc
	o.lSide = o.expression[:i]
	o.rSide = o.expression[i+1:]
}

func (e *eq) isdelimeter(cc string) (b bool) {
	if cc == e.delimeters[0] {
		b = true
		return
	}
	return
}
func (o *op) isdelimeter(cc string) (b bool) {
	if (cc == o.delimeters[0]) || (cc == o.delimeters[1]) {
		b = true
		return
	}
	return
}

func (o *op) lsplit(e *eq) {
	for i, c := range o.expression {
		cc := fmt.Sprintf("%c", c)
		if (cc == "+") || (cc == "-") {
			o.loperator = true
			o.delimeter = cc
			o.lSide = o.expression[:i]
			o.rSide = o.expression[i+1:]

		}
	}
}

func (o *op) rsplit(e *eq) {
	for i, c := range o.expression {
		cc := fmt.Sprintf("%c", c)
		if (cc == "+") || (cc == "-") {
			o.delimeter = cc
			o.lSide = o.expression[:i]
			o.rSide = o.expression[i+1:]

		}
	}
}
