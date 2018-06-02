package proceq
// this is going to be a funtion based progam
// first thing, grab each term.
// objects: term terms{[]terms}
// functions: getterm, stackterm
// term methods, isnegative, ispowerd, isvared, isvarpowered{if varered is false then this will not need to be run}, 
// isleftsideoeq{check the terms object to see if there is a equal during term processing to see if this is true, isleftsideofop{check the terms 
//	object to see if hasoperator is true} 
// identify the term by looking from the left of the expression, then finding , plus , minus, or equal 
//
import "testing"

func TestProcreq(t *testing.T) {
      t.Ruen("testing procreq using '2*x**2 + 3*y**2 = 15' \n", fun(t *testing.T) ) {
		  expression := "2*x**2 + 3*y**2 = 15"
		  proceq(expression)
	  } 
