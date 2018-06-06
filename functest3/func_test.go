package main

import "testing"

func Testfunc(t *testing.T) {
	t.Run("testing efind", func(t *testing.T) {
		gotlop, gothaslop := efind("2*x**2 + 3*y**2 = 15", "+")
		wantlop := "e"
		wanthaslop := false
		if gotlop != wantlop {
			t.Errorf("got '%s' want '%s'", gotlop, wantlop)
		}
		if gothaslop != wanthaslop {
			t.Errorf("got '%t' want '%t'", gothaslop, wanthaslop)
		}
		gotrop, gothasrop := efind("2*x**2 + 3*y**2 = 15", "+")
		wantrop := "+"
		wanthasrop := true
		if gotlop != wantlop {
			t.Errorf("got '%s' want '%s'", gotrop, wantrop)
		}
		if gothaslop != wanthaslop {
			t.Errorf("got '%t' want '%t'", gothasrop, wanthasrop)
		}
	})

}
