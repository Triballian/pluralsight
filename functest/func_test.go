package main

import (
	"testing"
)

func TestDelspace(t *testing.T) {
	t.Run("Deleteing space of ' 2*x**2 + 3*y**2 = 15\n'", func(t *testing.T) {
		var e eq
		e.expression = "2*x**2 + 3*y**2 = 15"
		e.delSpace()

		got := e.expression

		want := "2*x**2+3*y**2=15"
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}


}

func split(t *testing.T){
	t.Run("spliting the expression by equal sighn and operator", func(t *testing.T)){
		var e eq
		var o op
		e.expression = "2*x**2 + 3*y**2 = 15"

	}
}
