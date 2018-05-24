package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func prmptforInput() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nType the equation to be identified as either a function or not a function")
	fmt.Print("\n>>")

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(text)
}
