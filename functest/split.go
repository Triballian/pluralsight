package main

import "fmt"

func (e *eq) split(delimeter string) {
	for i, c := range e.expression {
		cc := fmt.Sprintf("%c", c)
		if cc == delimeter {
			fmt.Printf("There is an equal")
			e.delimeter = delimeter
			e.leftSide = e.expression[:i]
			e.rightSide = e.expression[i+1:]

		}
	}
}
