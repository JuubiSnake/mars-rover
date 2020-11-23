package travel

import "fmt"

// ParseDirectionError is an error that relates to an invalid Direction being parsed.
type ParseDirectionError struct {
	Direction string
}

// Error returns a message relating to the direction that was unable to be parsed.
func (p *ParseDirectionError) Error() string {
	return fmt.Sprintf("'%s' is not a valid direction", p.Direction)
}

// ParseMovementError is an error that relates to an invalid Movement being parsed.
type ParseMovementError struct {
	Move string
}

// Error returns a message relating to the movement that was unable to be parsed.
func (p *ParseMovementError) Error() string {
	return fmt.Sprintf("'%s' is not a valid move", p.Move)
}
