package main

import "fmt"

type terms struct {
	expression   string
	elside       string
	erside       string
	hasequals    bool
	hasloperator bool
	hasroperator bool
	loperator    string
	roperator    string
	//	terms        []term
}

//func (t *terms) getequals() {
// check expression for equal sign then parse elside an erside
// if there is an equl sign flag hasequals
//	fmt.Println("SOMETGING")

//}
// before you run this check to see if ther is an equals sign
// flag hasleftoperator
// dont forget to use string(b) where b is []byte to convert a series of characters
// like e[1:5] into a string

func (t *terms) getoperators(lside bool) {
	if lside {
		t.loperator, t.hasloperator = efind(t.elside, "+")
		t.roperator, t.hasroperator = efind(t.erside, "-")
		if t.roperator != "" {
			t.hasroperator = true
		}
	}

}

//search for item ignore the first character
func efind(t string, item string) (operator string, hasitem bool) {

	for sindex, c := range t {
		if sindex == 0 {
			continue
		}
		cstring := fmt.Sprintf("%c", c)
		if (cstring == "+") || (cstring == "-") {
			operator = cstring
			hasitem = true
			return
		}

	}
	operator = ""
	hasitem = false
	return

}

func main() {

}
