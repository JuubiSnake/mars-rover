package travel

import (
	"fmt"
	"testing"
)

func Test_combineDirectionWithMovement(t *testing.T) {
	d := North
	m := Left

	expected := fmt.Sprintf("%s-%s", d, m)
	if combineDirectionWithMovement(d, m) != expected {
		t.Fatalf("expected combineDirectionWithMovement to produce %s - instead got %s", expected, combineDirectionWithMovement(d, m))
	}
}

func Test_ParseDirection_InvalidLength(t *testing.T) {
	directionInvalid := "RR"
	_, err := ParseDirection(directionInvalid)
	if err == nil {
		t.Fatalf("ParseDirection should have failed with %s", directionInvalid)
	}
	pe, ok := err.(*ParseDirectionError)
	if !ok {
		t.Fatalf("ParseDirection should have produced ParseDirectionError - got %T instead", err)
	}
	if pe.Direction != directionInvalid {
		t.Fatalf("ParseDirectionError should have failed on direction %s - failed on %s instead", directionInvalid, pe.Direction)
	}
}

func Test_ParseDirection_InvalidDirection(t *testing.T) {
	directionInvalid := "R"
	_, err := ParseDirection(directionInvalid)
	if err == nil {
		t.Fatalf("ParseDirection should have failed with %s", directionInvalid)
	}
	pe, ok := err.(*ParseDirectionError)
	if !ok {
		t.Fatalf("ParseDirection should have produced ParseDirectionError - got %T instead", err)
	}
	if pe.Direction != directionInvalid {
		t.Fatalf("ParseDirectionError should have failed on direction %s - failed on %s instead", directionInvalid, pe.Direction)
	}
}

func testParseDirection(t *testing.T, direction string, expected Direction) {
	d, err := ParseDirection(direction)
	if err != nil {
		t.Fatalf("ParseDirection should not have failed - got the following error: %v", err)
	}
	if d != expected {
		t.Fatalf("expected %s to produce direction %s - got %s instead", direction, expected, d)
	}
}

func Test_ParseDirections(t *testing.T) {
	testParseDirection(t, "N", North)
	testParseDirection(t, "E", East)
	testParseDirection(t, "W", West)
	testParseDirection(t, "S", South)
}

func Test_ParseMovement_InvalidLength(t *testing.T) {
	invalidMovement := "RR"
	_, err := ParseMovement(invalidMovement)
	if err == nil {
		t.Fatalf("ParseMovement should have failed with %s", invalidMovement)
	}
	pe, ok := err.(*ParseMovementError)
	if !ok {
		t.Fatalf("ParseMovement should have produced ParseMovementError - got %T instead", err)
	}
	if pe.Move != invalidMovement {
		t.Fatalf("ParseMovementError should have failed on movement %s - failed on %s instead", invalidMovement, pe.Move)
	}
}

func Test_ParseMovement_InvalidMovement(t *testing.T) {
	invalidMovement := "E"
	_, err := ParseMovement(invalidMovement)
	if err == nil {
		t.Fatalf("ParseMovement should have failed with %s", invalidMovement)
	}
	pe, ok := err.(*ParseMovementError)
	if !ok {
		t.Fatalf("ParseMovement should have produced ParseMovementError - got %T instead", err)
	}
	if pe.Move != invalidMovement {
		t.Fatalf("ParseMovementError should have failed on movement %s - failed on %s instead", invalidMovement, pe.Move)
	}
}

func testParseMovement(t *testing.T, movement string, expected Movement) {
	m, err := ParseMovement(movement)
	if err != nil {
		t.Fatalf("ParseMovement should not have failed - got the following error: %v", err)
	}
	if m != expected {
		t.Fatalf("expected %s to produce movement %s - got %s instead", movement, expected, m)
	}
}

func Test_ParseMovements(t *testing.T) {
	testParseMovement(t, "L", Left)
	testParseMovement(t, "R", Right)
	testParseMovement(t, "M", Move)
}

type testDirectionalMovementOpts struct {
	Direction         Direction
	Movement          Movement
	ExpectedX         int
	ExpectedY         int
	ExpectedDirection Direction
}

func testDirectionalMovement(t *testing.T, opts *testDirectionalMovementOpts) {
	if opts == nil {
		t.Fatal("opts cannot be nil for testDirectionalMovement")
	}
	x, y, d := Travel(opts.Direction, opts.Movement)
	if x != opts.ExpectedX || y != opts.ExpectedY || opts.ExpectedDirection != d {
		t.Fatalf(
			"Travel should have returned x: %d y: %d d: %s for movement %s in direction %s - got x: %d y: %d d: %s instead",
			opts.ExpectedX,
			opts.ExpectedY,
			opts.ExpectedDirection,
			opts.Movement,
			opts.Direction,
			x,
			y,
			d,
		)
	}
}

func Test_Travel_North(t *testing.T) {
	opts := &testDirectionalMovementOpts{
		Direction:         North,
		Movement:          Left,
		ExpectedX:         0,
		ExpectedY:         0,
		ExpectedDirection: West,
	}
	testDirectionalMovement(t, opts)
	opts = &testDirectionalMovementOpts{
		Direction:         North,
		Movement:          Right,
		ExpectedX:         0,
		ExpectedY:         0,
		ExpectedDirection: East,
	}
	testDirectionalMovement(t, opts)
	opts = &testDirectionalMovementOpts{
		Direction:         North,
		Movement:          Move,
		ExpectedX:         0,
		ExpectedY:         1,
		ExpectedDirection: North,
	}
	testDirectionalMovement(t, opts)
}

func Test_Travel_East(t *testing.T) {
	opts := &testDirectionalMovementOpts{
		Direction:         East,
		Movement:          Left,
		ExpectedX:         0,
		ExpectedY:         0,
		ExpectedDirection: North,
	}
	testDirectionalMovement(t, opts)
	opts = &testDirectionalMovementOpts{
		Direction:         East,
		Movement:          Right,
		ExpectedX:         0,
		ExpectedY:         0,
		ExpectedDirection: South,
	}
	testDirectionalMovement(t, opts)
	opts = &testDirectionalMovementOpts{
		Direction:         East,
		Movement:          Move,
		ExpectedX:         1,
		ExpectedY:         0,
		ExpectedDirection: East,
	}
	testDirectionalMovement(t, opts)
}

func Test_Travel_West(t *testing.T) {
	opts := &testDirectionalMovementOpts{
		Direction:         West,
		Movement:          Left,
		ExpectedX:         0,
		ExpectedY:         0,
		ExpectedDirection: South,
	}
	testDirectionalMovement(t, opts)
	opts = &testDirectionalMovementOpts{
		Direction:         West,
		Movement:          Right,
		ExpectedX:         0,
		ExpectedY:         0,
		ExpectedDirection: North,
	}
	testDirectionalMovement(t, opts)
	opts = &testDirectionalMovementOpts{
		Direction:         West,
		Movement:          Move,
		ExpectedX:         -1,
		ExpectedY:         0,
		ExpectedDirection: West,
	}
	testDirectionalMovement(t, opts)
}

func Test_Travel_South(t *testing.T) {
	opts := &testDirectionalMovementOpts{
		Direction:         South,
		Movement:          Left,
		ExpectedX:         0,
		ExpectedY:         0,
		ExpectedDirection: East,
	}
	testDirectionalMovement(t, opts)
	opts = &testDirectionalMovementOpts{
		Direction:         South,
		Movement:          Right,
		ExpectedX:         0,
		ExpectedY:         0,
		ExpectedDirection: West,
	}
	testDirectionalMovement(t, opts)
	opts = &testDirectionalMovementOpts{
		Direction:         South,
		Movement:          Move,
		ExpectedX:         0,
		ExpectedY:         -1,
		ExpectedDirection: South,
	}
	testDirectionalMovement(t, opts)
}

func Test_Travel_Unknown(t *testing.T) {
	opts := &testDirectionalMovementOpts{
		Direction:         UnknownDirection,
		Movement:          UnknownMovement,
		ExpectedX:         0,
		ExpectedY:         0,
		ExpectedDirection: UnknownDirection,
	}
	testDirectionalMovement(t, opts)
}
