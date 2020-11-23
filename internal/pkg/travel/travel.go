package travel

import "fmt"

const (
	// North relates to the instruction 'N'.
	North Direction = "N"
	// East relates to the instruction 'E'.
	East Direction = "E"
	// West relates to the instruction 'W'.
	West Direction = "W"
	// South relates to the instruction 'S'.
	South Direction = "S"
	// UnknownDirection relates to an instruction that is unknown.
	UnknownDirection Direction = "_"

	// Left relates to the instruction 'L'.
	Left Movement = "L"
	// Right relates to the instruction 'R'.
	Right Movement = "R"
	// Move relates to the instruction 'M'.
	Move Movement = "M"
	// UnknownMovement relates to an instruction that is unknown.
	UnknownMovement Movement = "_"
)

var (
	northLeft  = combineDirectionWithMovement(North, Left)
	northRight = combineDirectionWithMovement(North, Right)
	northMove  = combineDirectionWithMovement(North, Move)

	eastLeft  = combineDirectionWithMovement(East, Left)
	eastRight = combineDirectionWithMovement(East, Right)
	eastMove  = combineDirectionWithMovement(East, Move)

	westLeft  = combineDirectionWithMovement(West, Left)
	westRight = combineDirectionWithMovement(West, Right)
	westMove  = combineDirectionWithMovement(West, Move)

	southLeft  = combineDirectionWithMovement(South, Left)
	southRight = combineDirectionWithMovement(South, Right)
	southMove  = combineDirectionWithMovement(South, Move)
)

// combineDirectionWithMovement is a simple helper function that maps a
// single Direction and Movement instruction together.
func combineDirectionWithMovement(d Direction, m Movement) string {
	return fmt.Sprintf("%s-%s", d, m)
}

// Direction is a type that relates to the instruction-set for directional headings
// There are 5 possible instructions; N, E, W, S, _.
type Direction string

// ParseDirection takes a strings and aliases it to a Direction instruction.
// It returns a ParseDirectionError if the string is greater than len 1 and
// is not part of the instruction-set consisting of N, E, W, S.
func ParseDirection(dir string) (Direction, error) {
	if len(dir) != 1 {
		return UnknownDirection, &ParseDirectionError{Direction: dir}
	}
	switch dir {
	case string(North):
		return North, nil
	case string(West):
		return West, nil
	case string(East):
		return East, nil
	case string(South):
		return South, nil
	default:
		return UnknownDirection, &ParseDirectionError{Direction: dir}
	}
}

// Movement is a type that relates to the instruction-set for movement.
// There are 4 possible instructions; L, R, M, _.
type Movement string

// ParseMovement takes a strings and aliases it to a Movement instruction.
// It returns a ParseMovementError if the string is greater than len 1 and
// is not part of the instruction-set consisting of L, R, M.
func ParseMovement(m string) (Movement, error) {
	if len(m) != 1 {
		return UnknownMovement, &ParseMovementError{Move: m}
	}
	switch m {
	case string(Left):
		return Left, nil
	case string(Right):
		return Right, nil
	case string(Move):
		return Move, nil
	default:
		return UnknownMovement, &ParseMovementError{Move: m}
	}
}

// Travel returns a co-ordinal vector based on a given direction and movement you wish
// to travel in.
// I.E on a 2D surface - if you wish to move left (L) when facing north (N) you will
// remain stationary (0, 0) but now face west (W).
func Travel(direction Direction, move Movement) (int, int, Direction) {
	directionalMovement := combineDirectionWithMovement(direction, move)
	switch directionalMovement {
	case northLeft:
		return 0, 0, West
	case northRight:
		return 0, 0, East
	case northMove:
		return 0, 1, North
	case eastLeft:
		return 0, 0, North
	case eastRight:
		return 0, 0, South
	case eastMove:
		return 1, 0, East
	case southLeft:
		return 0, 0, East
	case southRight:
		return 0, 0, West
	case southMove:
		return 0, -1, South
	case westLeft:
		return 0, 0, South
	case westRight:
		return 0, 0, North
	case westMove:
		return -1, 0, West
	default:
		return 0, 0, direction
	}
}
