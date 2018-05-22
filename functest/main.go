package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type eq struct {
	equation string
	noterms  int
	terms    []*term

	// coefficient terms or constant terms

}

type term struct {
	lside      bool
	operator   string // left side, right side of operator or there is no operator adjecent to this term.
	coeficient int    // every term has a coeffecient even if the coefficient is 1
	yvar       rune   // x, y, or nil

	exponent int // even if the exponent is zero
}

func prmptforInput() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nType the equation to be identified as either a function or not a function")
	fmt.Print("\n>>")

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func main() {
	var e eq
	var pe = &e
	pe.equation = prmptforInput()
	fmt.Printf("equation : %s", pe.equation)
	seq := strings.Split(pe.equation, `=`)
	proceq(pe)

	fmt.Printf("Left side is %s\n", seq[0])

}
