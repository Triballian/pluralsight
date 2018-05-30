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
	})

}

func TestSplit(t *testing.T) {
	t.Run("spliting the expression by equal sign", func(t *testing.T) {
		var e eq
		var o op
		e.expression = "2*x**2+3*y**2=15"
		o.expression = "2*x**2+3*y**2=15"
		split(&e)
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

		split(&o)
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

	})
}
