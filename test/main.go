package main

import "fmt"

type basetank struct {
	// spreadshot
	name    string
	power   int
	ttype   string
	color   string
	shape   string
	turrets int
	// size  int
}
type spreadshot struct {
	basetank
	// name  string
	// power int
	// ttype string
	size int
}

type tank interface {
	printname()
}

func main() {
	runapp()
}
func runapp() {
	var base basetank
	var spread spreadshot
	base.name = "base is base"
	base.power = 10
	base.ttype = "base is straight shot"
	base.color = "base is red"
	base.shape = "base is circle"
	base.turrets = 2
	spread.name = "spread is spread shot"
	spread.power = 40
	spread.ttype = "spread is bullet spammer"
	spread.size = 20
	spread.color = "spread is blue"
	spread.shape = "spread is square"
	spread.turrets = 3

	base.printname()
	spread.printname()
	spread.printtype()
	spread.printcolor()
	spread.printshape()
	//base.printshape() // spread inherits base's printcolor() but base does not inherit spread's printshape

	// printpower(spread)
	printturrets(base)
	printturrets(spread)
}

func (b basetank) printname() {
	fmt.Printf("%s\n", b.name)
}

func (b basetank) printcolor() {
	fmt.Printf("%s\n", b.color)
}

func (s spreadshot) printname() {
	fmt.Printf("%s\n", s.name)
}
func (s spreadshot) printtype() {
	fmt.Printf("%s\n", s.ttype)
}
func (s spreadshot) printshape() {
	fmt.Printf("%s\n", s.shape)
}

func printturrets(t tank) {
	t.printname()
}

// func printpower(b base) {
// 	fmt.Println("%d", b.power)

// }
