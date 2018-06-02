package main

type terms struct{
	expression string
	elside string
	erside string
	hasequals bool
	hasloperator bool
	hasroperator bool
	loperator string
	roperator string
	terms []term

}


func (t *terms) getequals {
	// check expression for equal sign then parse elside an erside
	// if there is an equl sign flag hasequals


}
// before you run this check to see if ther is an equals sign
// flag hasleftoperator
// dont for get to use string(b) where b is []byte to convert a series of characters
// like e[1:5] into a string
func (t *terms)getoperators(lside bool) {
	if lside{
		t.loperator, t.hasleftoperator = efind(t.elside, "+")
		t.roperator, t.hasroperator = efind(t.rlside, "-")
		if t.roperator != nil {
			t.hasroperator = true.
		}
	}
	
}
//search for item ignore the first character
func efind(t string, item string)(operator string, hasitem bool){
	var operator string
	var hasoperator bool
	for sindex, c =  range t {
		if sindex == 0 {
			continue
		}
		cstring = fmt.Sprintf("%c", c)
		if (cstring = "+") || (cstring = "-")
		operator = cstring
		hasitem = true
		return


	}
	operator = nil
	hasitem = false
	
}

func main() {

}