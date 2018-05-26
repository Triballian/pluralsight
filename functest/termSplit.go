package main

import "fmt"

// 1st find out if there is a var in the term

func (o *op) termSplit(h *head, e *eq) {
	var t term
	var v variable
	splitOnvar(h, &v, o, &t)
	if len(v.lside) > 0 {
		h.lsideOVar = true
		getbase(h, &v, o, &t)

	}
	if len(v.rside) > 0 {
		getExpo(&v, o, &t)
	}
}

func splitOnvar(h *head, v *variable, o *op, t *term) {
	if h.lsideOperator {
		for i, c := range o.lSide {
			cc := fmt.Sprintf("%c", c)
			if (cc == "x") || (cc == "y") {
				v.lside = o.lSide[:i]
				v.rside = o.lSide[i+1:]
				t.variable = cc
				t.lside = true

			}

		}
	} else {
		for i, c := range o.rSide {
			cc := fmt.Sprintf("%c", c)
			if (cc == "x") || (cc == "y") {
				v.rside = o.rSide[:i]
				v.rside = o.lSide[i+1:]
				t.variable = cc

			}
		}
	}
}

func getbase(h *head, v *variable, o *op, t *term) {
	for i, c := range v.lside {
		cc := fmt.Sprintf("%c", c)
		if cc == "*" {
			t.baseEx = o.lSide[:i]
			continue
		}
	}

}
func getExpo(v *variable, o *op, t *term) {
	for i, c := range v.rside {
		cc := fmt.Sprintf("%c", c)
		if cc == "*" {
			t.varEx = o.rSide[:i]
			continue
		}
	}

}
