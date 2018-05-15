package main

type fbq struct {
	equation    string
	operator    string
	isloperator bool
	lexp    []exp{} // is a slice for wich youcan grab the lentght len
	rexp    int // is a slice
	yexp bool
	// coefficient terms or constant terms

}
type exp struct {
	terms string
}

type cterm struct {
	varterm string 
	conterm string //  may be nil
}

type varterm struct {
	coeficient // every term has a coeffecient even if the coefficient is 1 
	yvar string // either x or y
}

type conterm struct {
	term string
}

type term struct {
	coefficient int
}

type variable struct {
	yvar bool
}



func main() {}
