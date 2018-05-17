package main

import (
	"fmt"

	"time"

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
	go salutations.Greet(greeting.CreatePrintFunction("<C>"), true, 6)
	salutations.Greet(greeting.CreatePrintFunction("?"), true, 6)
	// greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 6)
	time.Sleep(100 * time.Millisecond)
}
