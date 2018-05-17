package main

import (
	"fmt"

	"github.com/Triballian/pluralsight/hello/greeting"
)

func RenameToFrog(r greeting.Renameable) {
	r.Rename("Frog")
}

func main() {

	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
	}
	fmt.Fprintf(&salutations[0], "The count is %d", 10)
	done := make(chan bool)
	go func() {
		salutations.Greet(greeting.CreatePrintFunction("<C>"), true, 6)
		done <- true
	}()

	salutations.Greet(greeting.CreatePrintFunction("?"), true, 6)
	<-done //going to blok until it relceieve the varibale from the channle
	// greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 6)
}
