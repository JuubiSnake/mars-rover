package runner

import (
	"errors"
	"strconv"
	"testing"

	"github.com/juubisnake/mars-rover/internal/pkg/plateau"
	"github.com/juubisnake/mars-rover/internal/pkg/travel"
)

func testValidRun(t *testing.T, input, expected string) {
	actual, err := Run(input)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected runner to output:\n%s\ninstead got:\n%s", actual, expected)
	}
}

func TestRun_Example(t *testing.T) {
	input := `
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM`
	expected := "1 3 N\n5 1 E"
	testValidRun(t, input, expected)
}

func TestRun_Border(t *testing.T) {
	input := `
5 5
0 0 N
MMMMMRMMMMMRMMMMMRMMMMMR`
	expected := "0 0 N"
	testValidRun(t, input, expected)
}

func TestRun_Trim(t *testing.T) {
	input := `
 5 5       
    1 2 N
        LMLMLMLMM
    3 3 E
  MMRMMRMRRM`
	expected := "1 3 N\n5 1 E"
	testValidRun(t, input, expected)
}

func TestRun_MissingInput(t *testing.T) {
	input := `
5 5
1 3 N
`
	expectedLines := 2
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with incomplete input")
	}
	me, ok := err.(*MissingInputLinesError)
	if !ok {
		t.Fatalf("Run() should have produced a missingInputLinesError - got %T instead", err)
	}
	if me.Lines != expectedLines {
		t.Fatalf("missingInputLinesError should have detected %d lines within the input - got %d instead", expectedLines, me.Lines)
	}
}

func TestRun_EvenInputLength(t *testing.T) {
	input := `
5 5
1 3 N
LMLMRM
2 4 E
`
	expectedLines := 4
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with an even number of input lines")
	}
	ee, ok := err.(*EvenInputLinesError)
	if !ok {
		t.Fatalf("Run() should have produced a evenInputLinesError - got %T instead", err)
	}
	if ee.Lines != expectedLines {
		t.Fatalf("missingInputLinesError should have detected %d lines within the input - got %d instead", expectedLines, ee.Lines)
	}
}

func TestRun_InvalidSurfaceDimensions(t *testing.T) {
	input := `
5 5 3
1 3 N
LMLMRM
`
	expectedDimensions := 3
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with incorrect number of surface dimensions")
	}
	se, ok := err.(*SurfaceDimensionError)
	if !ok {
		t.Fatalf("Run() should have produced a surfaceDimensionError - got %T instead", err)
	}
	if se.Dimensions != expectedDimensions {
		t.Errorf("surfaceDimensionError should have detected %d dimensions within the input - got %d instead", expectedDimensions, se.Dimensions)
	}
}

func TestRun_buildSurface_unparseableXCoordinate(t *testing.T) {
	input := `
E 3
1 3 N
LMLMRM
`
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with invalid surface boundaries")
	}
	sb, ok := err.(*ParseSurfaceBoundaryError)
	if !ok {
		t.Fatalf("Run() should have produced a ParseSurfaceBoundaryError - got %T instead", err)
	}
	var ne *strconv.NumError
	if !errors.As(sb, &ne) {
		t.Fatal("ParseSurfaceBoundaryError should have a wrapped NumError")
	}
	if sb.Coordinate != "x" {
		t.Fatalf("ParseSurfaceBoundaryError should have failed on the x cordinate - failed on %s instead", sb.Coordinate)
	}
}

func TestRun_buildSurface_unparseableYCoordinate(t *testing.T) {
	input := `
3 E
1 3 N
LMLMRM
`
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with invalid surface boundaries")
	}
	sb, ok := err.(*ParseSurfaceBoundaryError)
	if !ok {
		t.Fatalf("Run() should have produced a ParseSurfaceBoundaryError - got %T instead", err)
	}
	var ne *strconv.NumError
	if !errors.As(sb, &ne) {
		t.Fatal("ParseSurfaceBoundaryError should have a wrapped NumError")
	}
	if sb.Coordinate != "y" {
		t.Fatalf("ParseSurfaceBoundaryError should have failed on the y cordinate - failed on %s instead", sb.Coordinate)
	}
}

func TestRun_buildSurface_invalidXBoundary(t *testing.T) {
	input := `
-1 0
1 3 N
LMLMRM
`
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with invalid x coordinate")
	}
	sb, ok := err.(*SurfaceError)
	if !ok {
		t.Fatalf("Run() should have produced a SurfaceError - got %T instead", err)
	}
	var ne *plateau.UpperBoundsError
	if !errors.As(sb, &ne) {
		t.Fatal("SurfaceError should have a wrapped OutOfBoundsError")
	}
	if ne.Coordinate != "x" {
		t.Fatalf("OutOfBoundsError should have failed on the x cordinate - failed on %s instead", ne.Coordinate)
	}
}

func TestRun_buildSurface_invalidYBoundary(t *testing.T) {
	input := `
0 -1
1 3 N
LMLMRM
`
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with invalid x coordinate")
	}
	sb, ok := err.(*SurfaceError)
	if !ok {
		t.Fatalf("Run() should have produced a SurfaceError - got %T instead", err)
	}
	var ne *plateau.UpperBoundsError
	if !errors.As(sb, &ne) {
		t.Fatal("SurfaceError should have a wrapped OutOfBoundsError")
	}
	if ne.Coordinate != "y" {
		t.Fatalf("OutOfBoundsError should have failed on the y cordinate - failed on %s instead", ne.Coordinate)
	}
}

func TestRun_BuildRobot_InvalidInstructionLength(t *testing.T) {
	input := `
5 5
1 3 N
LMLMLM
1 3 N F
LMLMRM
`
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with invalid robot instruction length")
	}
	_, ok := err.(*RobotInstructionLengthError)
	if !ok {
		t.Fatalf("Run() should have produced a RobotInstructionLengthError - got %T instead", err)
	}
}

func TestRun_BuildRobot_UnparseableXCoordinate(t *testing.T) {
	input := `
5 5
1 3 N
LMLMLM
B 2 N
LMLMRM
`
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with invalid robot starting x coordinate")
	}
	pe, ok := err.(*ParseRobotCoordinateError)
	if !ok {
		t.Fatalf("Run() should have produced a ParseRobotCoordinateError - got %T instead", err)
	}
	if pe.Coordinate != "x" || pe.Position != "B" {
		t.Fatalf("ParseRobotCoordinateError should have failed parsing x coordinate B: %v", pe)
	}
}

func TestRun_BuildRobot_UnparseableYCoordinate(t *testing.T) {
	input := `
5 5
1 3 N
LMLMLM
1 Y N
LMLMRM
`
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with invalid robot starting y coordinate")
	}
	pe, ok := err.(*ParseRobotCoordinateError)
	if !ok {
		t.Fatalf("Run() should have produced a ParseRobotCoordinateError - got %T instead", err)
	}
	if pe.Coordinate != "y" || pe.Position != "Y" {
		t.Fatalf("ParseRobotCoordinateError should have failed parsing y-coordinate Y: %v", pe)
	}
}

func TestRun_BuildRobot_UnparseableDirection(t *testing.T) {
	input := `
5 5
1 3 N
LMLMLM
1 3 F
LMLMRM
`
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with invalid robot starting direction")
	}
	de, ok := err.(*ParseRobotDirectionError)
	if !ok {
		t.Fatalf("Run() should have produced a ParseRobotDirectionError - got %T instead", err)
	}
	var wrapped *travel.ParseDirectionError
	if !(errors.As(de, &wrapped)) {
		t.Fatal("ParseRobotDirectionError should have wrapped a travel.ParseRobotDirectionError")
	}
	if wrapped.Direction != "F" {
		t.Fatalf("ParseRobotDirectionError should have failed to parse direction F - failed with %s", wrapped.Direction)
	}
}

func TestRun_BuildRobot_OutOfBounds(t *testing.T) {
	input := `
5 5
1 6 N
LMLMLM
1 3 E
LMLMRM
`
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with a robot starting position of out-of-bounds")
	}
	be, ok := err.(*RobotOutOfBoundsError)
	if !ok {
		t.Fatalf("Run() should have produced a RobotOutOfBoundsError - got %T instead", err)
	}
	if be.Y != 6 {
		t.Fatalf("robot should have been marked out of bounds with starting position of 1 6 - instead got: %v", be)
	}
}

func TestRun_GuideRobot_InvalidMovement(t *testing.T) {
	input := `
5 5
1 3 N
LMLMLMER
1 3 E
LMLMRM
`
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with a invalid movement")
	}
	pe, ok := err.(*ParseRobotMovementError)
	if !ok {
		t.Fatalf("Run() should have produced a ParseRobotMovementError - got %T instead", err)
	}
	var wrapped *travel.ParseMovementError
	if !(errors.As(pe, &wrapped)) {
		t.Fatal("ParseRobotMovementError should have wrapped a travel.ParseMovementError")
	}
	if wrapped.Move != "E" {
		t.Fatalf("ParseRobotMovementError should have errored on movement E - failed on %s instead", wrapped.Move)
	}
}

func TestRun_GuideRobot_OutOfBounds(t *testing.T) {
	input := `
5 5
1 1 N
MMMMMMM
1 3 E
LMLMRM
`
	_, err := Run(input)
	if err == nil {
		t.Fatal("Run() should have failed with a robot movement that is out-of-bounds")
	}
	pe, ok := err.(*RobotOutOfBoundsError)
	if !ok {
		t.Fatalf("Run() should have produced a RobotOutOfBoundsError - got %T instead", err)
	}
	if pe.Y != 6 {
		t.Fatalf("robot should have out-of-bounded on y coordinate 6 - got %d instead", pe.Y)
	}
}
