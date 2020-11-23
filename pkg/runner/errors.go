package runner

import "fmt"

// MissingInputLinesError is an error that is used whenever the number of instructions given
// to the Run command is less than the expected amount.
type MissingInputLinesError struct{ Lines int }

// Error outputs a message relating to the missing input.
func (m *MissingInputLinesError) Error() string {
	return fmt.Sprintf("the input should have at least %d lines - %d lines were detected", minimumInputLines, m.Lines)
}

// EvenInputLinesError is an error that is used whenever the input to the Run command has an even number
// of instructions.
// You can never have an even number of commands since a surface needs one line of instructions, while a robot
// requires two. This can be reduced to the rule 2n + 1, which is also the rule for an odd number.
type EvenInputLinesError struct{ Lines int }

// Error outputs a message relating to the even number of instructions within the input.
func (e *EvenInputLinesError) Error() string {
	return fmt.Sprintf("the input to this runner should have an odd number of lines - %d lines were detected", e.Lines)
}

// SurfaceDimensionError is an error that is thrown whenever a surface does not have two dimensions.
// I.E the instruction-set 2 2 3 is incorrect since that would relate to a three-dimensional coordinate.
type SurfaceDimensionError struct {
	Dimensions int
	Surface    string
}

// Error outputs a message relating to an incorrect number of dimensions for a surface.
func (s *SurfaceDimensionError) Error() string {
	return fmt.Sprintf("surface '%s' does not have the required number of surface dimensions - expected %d dimensions - detected %d instead", s.Surface, requiredSurfaceDimensions, s.Dimensions)
}

// ParseSurfaceBoundaryError is an error that is used whenever a coordinate for a surface cannot be parsed
// into an int.
type ParseSurfaceBoundaryError struct {
	Coordinate string
	Bounary    string
	Err        error
}

// Error outputs a message about the coordinate that was unable to be parsed.
func (s *ParseSurfaceBoundaryError) Error() string {
	return fmt.Sprintf("%s coordinate boundary '%s' cannot be transformed into an int", s.Coordinate, s.Bounary)
}

// Unwrap returns the error that is contained within the ParseSurfaceBoundaryError.
func (s *ParseSurfaceBoundaryError) Unwrap() error {
	return s.Err
}

// SurfaceError is an error that is returned whenever a surface is unable to be constructed.
// It wraps any error returned from plateau.New.
type SurfaceError struct {
	Err error
}

// Error outputs a simple message stating the surface could not be created.
func (s *SurfaceError) Error() string {
	return fmt.Sprintf("unable to create new surface")
}

// Unwrap returns the error that is contained within SurfaceError.
func (s *SurfaceError) Unwrap() error {
	return s.Err
}

// RobotInstructionLengthError is an error that is thrown whenever the number of instructions
// used to construct a robot is not equal to runner.robotInstructionLength
type RobotInstructionLengthError struct {
	Instructions string
	ID           int
}

// Error returns a message around the unexpected instruction length.
func (r *RobotInstructionLengthError) Error() string {
	return fmt.Sprintf("'%s' for robot ID %d does not have the required number of instructions - expected %d - detected %d instead", r.Instructions, r.ID, len(r.Instructions), robotInstructionLength)
}

// ParseRobotCoordinateError is an error that is thrown whenever a coordinate used
// for a robot from a instruction-set is incorrect.
type ParseRobotCoordinateError struct {
	ID         int
	Coordinate string
	Position   string
	Err        error
}

// Error outputs a message that relates to the coordinate that is unable to be parsed.
func (s *ParseRobotCoordinateError) Error() string {
	return fmt.Sprintf("robot ID %ds %s-coordinate '%s' cannot be transformed into an int", s.ID, s.Coordinate, s.Position)
}

// Unwrap returns the error that is contained within ParseRobotCoordinateError
func (s *ParseRobotCoordinateError) Unwrap() error {
	return s.Err
}

// ParseRobotDirectionError is an error that is thrown whenever a direction used
// for a robot is unable to be parsed.
type ParseRobotDirectionError struct {
	Direction string
	ID        int
	Err       error
}

// Error outputs a message that relates to the direction that is unable to be parsed.
func (p *ParseRobotDirectionError) Error() string {
	return fmt.Sprintf("unable to parse robot ID %ds direction: %s", p.ID, p.Direction)
}

// Unwrap returns the error contained within ParseRobotDirectionError.
func (p *ParseRobotDirectionError) Unwrap() error {
	return p.Err
}

// RobotOutOfBoundsError is an error that is returned whenever a robot moves out of bounds
// within a surface.
type RobotOutOfBoundsError struct {
	ID int
	X  int
	Y  int
}

// Error outputs a message that relates to the out-of-bound position.
func (r *RobotOutOfBoundsError) Error() string {
	return fmt.Sprintf("robot ID %d has moved out of bounds - X: %d Y: %d", r.ID, r.X, r.Y)
}

// ParseRobotMovementError is an error that is returned whenever a movement instruction is
// unable to be parsed.
type ParseRobotMovementError struct {
	Movement string
	ID       int
	Err      error
}

// Error outputs a message relating to the movement that is unable to be parsed.
func (p *ParseRobotMovementError) Error() string {
	return fmt.Sprintf("unable to parse robot ID %ds movement: %s", p.ID, p.Err)
}

// Unwrap returns the error that is contained within the ParseRobotMovementError.
func (p *ParseRobotMovementError) Unwrap() error {
	return p.Err
}
