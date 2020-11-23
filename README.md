# Mars Rover
[![Go Report Card](https://goreportcard.com/badge/github.com/juubisnake/mars-rover)](https://goreportcard.com/report/github.com/juubisnake/mars-rover)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/juubisnake/mars-rover)](https://pkg.go.dev/github.com/juubisnake/mars-rover)

A simple golang runner that allows you to simulate the movement of a series of robots against a plateau/surface via a instruction-set.

## Requirements

You will need:
- [`Golang` 1.15+](https://golang.org/doc/install)

## Overview

A squad of robotic rovers are to be landed by NASA on a plateau on Mars.

This plateau, which is curiously rectangular, must be navigated by the rovers so that their on board cameras can get a complete view of the surrounding terrain to send back to Earth.

A rover's position is represented by a combination of an x and y co-ordinates and a letter representing one of the four cardinal compass points.

The plateau is divided up into a grid to simplify navigation. An example position might be 0, 0, N, which means the rover is in the bottom left corner and facing North.

In order to control a rover, NASA sends a simple string of letters.

The possible letters are 'L', 'R' and 'M'. 'L' and 'R' makes the rover spin 90 degrees left or right respectively, without moving from its current spot.

'M' means move forward one grid point, and maintain the same heading.

Assume that the square directly North from (x, y) is (x, y+1).

## Instruction Format

The first line of input is the upper-right coordinates of the plateau, the lower-left coordinates are assumed to be 0,0.

The rest of the input is information pertaining to the rovers that have been deployed. Each rover has two lines of input.

The first line gives the rover's position, and the second line is a series of instructions telling the rover how to explore the plateau.

The position is made up of two integers and a letter separated by spaces, corresponding to the x and y co-ordinates and the rover's orientation.

Each rover will be finished sequentially, which means that the second rover won't start to move until the first one has finished moving.

Wrapping all of this up means the input is structured as follows:
```
5 5 // Creating a surface
1 2 N // Creating and placing a robot onto the surface
LMLMLMLMM // Moving the robot around the surface
3 3 E // Creating and placing a robot onto the surface
MMRMMRMRRM // Moving the robot around the surface
```
It is important to note that any whitespace surrounding the lines of the instruction-set will be stripped, so lines that proceed and trail with whitespace are still considered valid, for example:
```
        5 5
    1 2 N
LMLMM
```

However if there is whitespace between instructions within the instruction-set that are not syntatically correct - it will still be considered _invalid_, for example:
```
5   5
1  2 N
L  MLM MLM
```

## Usage

In order to run an instruction-set against the `mars-rover` package, you must use the `Run()` function within the `runner` package; for example:
```go
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
0 0 N
MMMMMRMMMMM
`
	result, err := runner.Run(instructions)
	if err != nil {
		log.Fatalf("failed while running instructions: %v", err)
	}
	fmt.Println(result)
}
```

You can find more examples within the [main file](./cmd/mars-rover/main.go).

If you wish to run these examples, simply run the following:
```shell
$ go run ./cmd/mars-rover/main.go
```
## Tests

This package comes a fleet of tests designed to ensure that simulator works with as much confidence as possible.

To run these tests simply run the following command:
```shell
$ go test ./...
```

## Formatting and Linting

To ensure the `.go` files are formatted and linted correctly, please run the following command:
```shell
$ go fmt ./... 
```
This will write any suggested changes the `gofmt` tool reports to the `.go` files.

## Future

In the future, the following will be implemented:
- Support for allowing the instruction-set to read from a `bufio.Reader` - this will allow us to stream an input from a large instruction-set and output the results per-line.
- Support for rewinding the robots positional history given they have fallen to an error - this could be part of an struct that errors within the robot package can contain.
