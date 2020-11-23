package runner

import (
	"strconv"
	"strings"

	"github.com/juubisnake/mars-rover/internal/pkg/plateau"
	"github.com/juubisnake/mars-rover/internal/pkg/robot"
	"github.com/juubisnake/mars-rover/internal/pkg/travel"
)

const (
	// minimumInputLines is a smallest input set possible for a simulation.
	minimumInputLines = 3
	// requiredSurfaceDimensions maps to the possible number of dimensions within a surface
	// E.G 1 is a 1D surface, 2 is a 2D surface, 3 is a 3D surface and so on.
	requiredSurfaceDimensions = 2
	// robotInstructionLength is the number of instructions required to create a robot.
	robotInstructionLength = 3
)

// manager contains helper functions that create and guides a
// robot along a given surface.
type manager struct {
	surface *plateau.Surface
	robot   *robot.Robot
}

// Run takes an instruction-set and uses it to generate a surface and
// guide a set a robots along its surface.
//
// Lines for the instruction-set must be delimited by '\n' and contain
// at least three lines in the following format:
//
//
// 5 5 <-- Line 0: This constructs a 2D surface.
//
// 1 1 N <-- Line 1: This constructs and places a robot on the surface.
//
// LMLLMLM <-- Line 2: This moves the robot along the surface.
//
//
// The input can repeat the instructions outlined on line 1 and 2 in order
// to place more robots on the surface, for example:
//
//
// 5 5
//
// 1 1 N
//
// LMLMLM
//
// 2 2 S
//
// LLLLLL
//
//
// Run will return the results of robots final resting positions in the following
// format:
//
//
// 2 2 E
//
// 1 1 N
//
//
// Note that output is delimited by '\n' and the number of lines is equal to
// the number of robots constructed.
//
// If an error occurs, Run will return the resting places of ALL robots that
// have successfully moved across the surface AND the error, for example:
//
//
// 5 5
//
// 1 1 N
//
// L
//
// 1 1 N
//
// MMMMMMM
//
//
// Will return a string "1 1 W" and an out-of-bounds error.
func Run(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) < minimumInputLines {
		return "", &MissingInputLinesError{Lines: len(lines)}
	}
	if len(lines)%2 == 0 {
		return "", &EvenInputLinesError{Lines: len(lines)}
	}
	surface, err := buildSurface(lines[0])
	if err != nil {
		return "", err
	}
	instructions := lines[1:]
	m := &manager{surface: surface}
	var output []string
	for i := 0; i+1 < len(instructions); i += 2 {
		robot, err := m.BuildRobot(i, instructions[i])
		if err != nil {
			return strings.Join(output, "\n"), err
		}
		m.robot = robot
		position, err := m.GuideRobot(instructions[i+1])
		if err != nil {
			return strings.Join(output, "\n"), err
		}
		output = append(output, position)
	}
	return strings.Join(output, "\n"), nil
}

// buildSurface constructs a plateau.Surface instance given a valid instruction.
func buildSurface(s string) (*plateau.Surface, error) {
	bounds := strings.Split(strings.TrimSpace(s), " ")
	if len(bounds) != requiredSurfaceDimensions {
		return nil, &SurfaceDimensionError{Dimensions: len(bounds), Surface: s}
	}
	boundX, err := strconv.Atoi(bounds[0])
	if err != nil {
		return nil, &ParseSurfaceBoundaryError{Coordinate: "x", Bounary: bounds[0], Err: err}
	}
	boundY, err := strconv.Atoi(bounds[1])
	if err != nil {
		return nil, &ParseSurfaceBoundaryError{Coordinate: "y", Bounary: bounds[1], Err: err}
	}
	surface, err := plateau.New(boundX, boundY)
	if err != nil {
		return nil, &SurfaceError{Err: err}
	}
	return surface, nil
}

// BuildRobot constructs a robot given a valid instruction.
func (m *manager) BuildRobot(id int, s string) (*robot.Robot, error) {
	config := strings.Split(strings.TrimSpace(s), " ")
	if len(config) != robotInstructionLength {
		return nil, &RobotInstructionLengthError{Instructions: s, ID: id}
	}
	x, err := strconv.Atoi(config[0])
	if err != nil {
		return nil, &ParseRobotCoordinateError{ID: id, Coordinate: "x", Position: config[0], Err: err}
	}
	y, err := strconv.Atoi(config[1])
	if err != nil {
		return nil, &ParseRobotCoordinateError{ID: id, Coordinate: "y", Position: config[1], Err: err}
	}
	direction, err := travel.ParseDirection(config[2])
	if err != nil {
		return nil, &ParseRobotDirectionError{Direction: config[2], ID: id, Err: err}
	}
	if m.surface.IsOutOfBounds(x, y) {
		return nil, &RobotOutOfBoundsError{ID: id, X: x, Y: y}
	}
	return robot.New(id, x, y, direction), nil
}

// GuideRobot guides a robot around a surface given a valid instruction.
func (m *manager) GuideRobot(commands string) (string, error) {
	fmtdCommands := strings.TrimSpace(commands)
	for i := range fmtdCommands {
		cmd := string(fmtdCommands[i])
		move, err := travel.ParseMovement(cmd)
		if err != nil {
			return "", &ParseRobotMovementError{Movement: cmd, Err: err, ID: m.robot.GetID()}
		}
		m.robot.Move(travel.Travel(m.robot.GetDirection(), move))
		if m.surface.IsOutOfBounds(m.robot.GetX(), m.robot.GetY()) {
			return "", &RobotOutOfBoundsError{ID: m.robot.GetID(), X: m.robot.GetX(), Y: m.robot.GetY()}
		}
	}
	return m.robot.String(), nil
}
