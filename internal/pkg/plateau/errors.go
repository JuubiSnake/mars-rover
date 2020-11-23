package plateau

import "fmt"

// UpperBoundsError is an error that is returned whenever the plateau surface's upper boundaries
// are less than the lower boundaries defined within the plateau package; see plateau.lowerBoundX
// and plateau.lowerBoundY
type UpperBoundsError struct {
	Coordinate string
	Value      int
	LowerBound int
}

// Error returns a message containing which coordinate has been found to be invalid as an upper bound.
func (o *UpperBoundsError) Error() string {
	return fmt.Sprintf("the upper bound %d for coordinate %s must be greater than or equal to %d", o.Value, o.Coordinate, o.LowerBound)
}
