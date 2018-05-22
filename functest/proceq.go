package main

import (
	"fmt"
)

func proceq(e *eq) {
	t := 0
	for i, c := range e.equation {
		cc := fmt.Sprintf("%c", c)

		if cc == "\n" {
			continue
		}

		if (cc == "+") || (cc == "-") {
			t++

			createTerm(e, i, t, cc)
		}
	}
}

func createTerm(e *eq, i int, t int, c string) {
	var term term
	pt := &term
	pt.operator = c
	if e.terms == nil {
		fmt.Printf("e.terms is %v", e.terms)
	}
	e.terms = append(e.terms, pt)

}
