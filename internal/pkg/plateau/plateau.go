package plateau

import "fmt"

const (
	// lowerBoundX is the lower-left x-coordinate boundary for a plateau's surface.
	lowerBoundX = 0
	// lowerBoundY is the lower-left y-coordinate boundary for a plateau's surface.
	lowerBoundY = 0
)

// Surface is a representation of a plateau upon which objects can traverse.
type Surface struct {
	LowerBoundX int
	LowerBoundY int
	UpperBoundX int
	UpperBoundY int
}

// New creates a new Surface with a given upper-right boundary, represented via x and y coordinates,
// with a lower-left boundary of 0,0.
// This function will error if any of the upper-right boundaries are less than the lower-left ones.
func New(upperBoundX, upperBoundY int) (*Surface, error) {
	if upperBoundX < lowerBoundX {
		return nil, &UpperBoundsError{Coordinate: "x", Value: upperBoundX, LowerBound: lowerBoundX}
	}
	if upperBoundY < lowerBoundY {
		return nil, &UpperBoundsError{Coordinate: "y", Value: upperBoundY, LowerBound: lowerBoundY}
	}
	return &Surface{
		LowerBoundX: lowerBoundX,
		LowerBoundY: lowerBoundY,
		UpperBoundX: upperBoundX,
		UpperBoundY: upperBoundY,
	}, nil
}

// IsOutOfBounds checks if a coordinate is out of bounds within a given surface.
func (s *Surface) IsOutOfBounds(x, y int) bool {
	return checkOutOfBounds(x, s.UpperBoundX, s.LowerBoundX) || checkOutOfBounds(y, s.UpperBoundY, s.LowerBoundY)
}

// String outputs a simple representation of a given surface.
func (s *Surface) String() string {
	return fmt.Sprintf("Surface | lower-bounds [%d,%d] - upper-bounds [%d,%d]", s.LowerBoundX, s.LowerBoundY, s.UpperBoundX, s.UpperBoundY)
}

// checkOutOfBounds checks if a given value is less than a lower-bound and greater than a higher
// bound.
func checkOutOfBounds(p, upper, lower int) bool {
	return p < lower || p > upper
}
