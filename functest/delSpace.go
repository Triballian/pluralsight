package main

import (
	"fmt"
	"unicode"
)

func (e *eq) delSpace() {
	var f string
	var fstring string
	// remove space and convert text into a slice
	for _, value := range e.expression {
		if !unicode.IsSpace(value) {
			fstring = fmt.Sprintf("%c", value)
			f = f + fstring
		}

	}
	e.expression = f

}
