package main

import (
	"fmt"
	"log"

	"github.com/juubisnake/mars-rover/pkg/runner"
)

// main is an example of how to use the mars-rover runner.
func main() {
	// a typical example of running a single robot across a surface.
	instructions := `
5 5
1 2 N
LMLMLMLMM
`
	result, err := runner.Run(instructions)
	if err != nil {
		log.Fatalf("failed while running instructions: %v", err)
	}
	fmt.Println(result)

	// a example of running a multiple rovers across a surface.
	instructions = `
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
`
	result, err = runner.Run(instructions)
	if err != nil {
		log.Fatalf("failed while running instructions: %v", err)
	}
	fmt.Println(result)

	// a example of outputting the first rovers resting place given
	// the second rover fails.
	instructions = `
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRMMMMMMMMMM
`
	result, err = runner.Run(instructions)
	if err == nil {
		log.Fatal("instruction should have failed")
	}
	fmt.Println(result)
}
