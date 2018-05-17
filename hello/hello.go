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
	c2 := make(chan greeting.Salutation)
	// call a goroutine that will that will fill the channel
	go salutations.ChannelGreeter(c)
	go salutations.ChannelGreeter(c2)
	for {
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
			} else {
				return
			}
		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}
		default:
			fmt.Println("Waiting...")

		}

	}
}
