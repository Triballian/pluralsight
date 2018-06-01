package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	expression := "22*x**2 + 3*y**2 = 15"
	lenex := len(expression)
	proceq(expression, lenex)
}

func proceq(e string, elen int) {
	// var ibuff []int

	for eindex, cc := range e {
		var estring string
		var tstring string
		var tstart int

		estring = fmt.Sprintf("%c", cc)
		_, error := strconv.Atoi(fmt.Sprintf("%v", estring))
		if error == nil {
			tstart = eindex
			for tindex, ccc := range e[eindex:] {

				tstring = fmt.Sprintf("%c", ccc)
				_, error = strconv.Atoi(fmt.Sprintf("%v", tstring))

				if error != nil {
					fmt.Printf("tstart is %v: tindex-1 is %v\n", tstart, tindex-1)
					if tstart == tindex-1 {
						//ascinumber, _ := strconv.Atoi(string(fmt.Sprintf("%c", e[tindex-1])))
						temn, _ := fmt.Printf("%v", e[:1])
						// ascinumber, _ := strconv.Atoi(fmt.Printf("%v", e[tindex-1]))
						fmt.Printf("temn is %v:\n tindex is %v\n", temn, tindex)
					} else {
						ascinumber, _ := strconv.Atoi(fmt.Sprintf("%c", e[tstart:tindex-1]))
						fmt.Printf("ascinumber is %v:\n tindex is %v\n", ascinumber, tindex)
					}

					temt, _ := fmt.Printf("%v", e[:2])
					fmt.Printf("temt is %v\n", temt)
					fmt.Printf("%v", e[0])
					fmt.Print("Press 'Enter' to continue...")
					bufio.NewReader(os.Stdin).ReadBytes('\n')
				}
			}
		}
	}

}
