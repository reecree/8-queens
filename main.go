package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/reecree/8-queens/src/board"
)

// A main function that reads command line inputs and calls appropriate
// mix tape function
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Improper Usage: 1 arguments required")
		return
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%s is not a number. Try again\n", args[0])
		return
	}

	board.RunHillClimbing(i, true)
}
