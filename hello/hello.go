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
	c := make(chan greeting.Salutation)
	// call a goroutine that will that will fill the channel
	go salutations.ChannelGreeter(c)
	for s := range c {
		fmt.Println(s.Name)
	}
}
