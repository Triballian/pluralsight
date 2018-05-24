package main

import "fmt"

func (e *eq) split(delimeter string) {
	for i, c := range e.header.expression {
		cc := fmt.Sprintf("%c", c)
		if cc == delimeter {
			fmt.Printf("There is an %s", cc)
			e.delimeter = delimeter
			e.leftSide = e.header.expression[:i]
			e.rightSide = e.header.expression[i+1:]

		}
	}
}
