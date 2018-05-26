package main

import "testing"

func TestAll(t *testing.T) {
	is = is.New(t)
	// var v variable

	var tests = []struct {
		Fn func(v *variable, o *op, t *term)
		S  string
	}{
		{Fn: func(v *variable, o *op, t *term) {
			getExpo(v, o, t)
		},
			S: "2",
		},
	}

	for _, test := range tests {
		test.Fn(v, o, t)
		is.Equal(test.S, t.varEx)
	}

}
