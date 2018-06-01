package main

import (
	"testing"
)

// every expression has two sides of an equal even if on side is nil
// every equation side has two sides of an operator even if one side is nil
// every operator side has two sides of a varibale, even if there no variables on that side. it will be nil variable.
// if it is a variable and the left side is a nil it has a coefficent of 1. if the right side is nil it has an exponent of 0
// if the left side of the equations leftside of the operators variable is nil then it is a constant. if the constant has a right side then it is a constant wit a exponent.
// all these objects fit inside each other.
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
	})

}

func TestSplit(t *testing.T) {
	t.Run("spliting the expression by equal sign", func(t *testing.T) {
		var e eq
		var o op
		var lo op
		var ro op

		e.expression = "2*x**2+3*y**2=15"
		o.expression = "2*x**2+3*y**2"
		e.delimeters = append(e.delimeters, "=")

		o.delimeters = append(o.delimeters, "+", "-")
		lo.delimeters = append(lo.delimeters, "x", "y")

		ro.delimeters = append(lo.delimeters, "x", "y")
		split(&e, e.expression)
		egotl := e.lSide
		egotr := e.rSide

		ewantl := "2*x**2+3*y**2"
		ewantr := "15"
		if egotl != ewantl {
			t.Errorf("gotl '%s' wantl '%s'", egotl, ewantl)
		}
		if egotr != ewantr {
			t.Errorf("gotr '%s' wantr '%s'", egotr, ewantr)
		}

		split(&o, e.lSide)
		ogotl := o.lSide
		ogotr := o.rSide

		owantl := "2*x**2"
		owantr := "3*y**2"
		if ogotl != owantl {
			t.Errorf("gotl '%s' wantl '%s'", ogotl, owantl)
		}
		if ogotr != owantr {
			t.Errorf("gotr '%s' wantr '%s'", ogotr, owantr)
		}

		split(&lo, "2*x**2")
		logotl := lo.lSide
		logotr := lo.rSide

		lowantl := "2*x"
		lowantr := "**2"
		if logotl != lowantl {
			t.Errorf("logotl '%s' wantl '%s'", logotl, lowantl)
		}
		if logotr != lowantr {
			t.Errorf("logotr '%s' wantr '%s'", logotr, lowantr)
		}

	})
}
