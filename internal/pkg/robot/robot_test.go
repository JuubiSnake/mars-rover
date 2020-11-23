package robot

import (
	"fmt"
	"testing"

	"github.com/juubisnake/mars-rover/internal/pkg/travel"
)

func Test_GetID(t *testing.T) {
	id, x, y, d := 0, 1, 1, travel.East
	r := New(id, x, y, d)
	if r.GetID() != id {
		t.Fatalf("expected GetID to be %d - got %d instead", id, r.GetID())
	}
}

func Test_GetX(t *testing.T) {
	id, x, y, d := 0, 1, 1, travel.East
	r := New(id, x, y, d)
	if r.GetX() != x {
		t.Fatalf("expected GetX to be %d - got %d instead", x, r.GetX())
	}
}

func Test_GetY(t *testing.T) {
	id, x, y, d := 0, 1, 1, travel.East
	r := New(id, x, y, d)
	if r.GetY() != y {
		t.Fatalf("expected GetY to be %d - got %d instead", y, r.GetY())
	}
}

func Test_GetDirection(t *testing.T) {
	id, x, y, d := 0, 1, 1, travel.East
	r := New(id, x, y, d)
	if r.GetDirection() != d {
		t.Fatalf("expected GetDirection to be %s - got %s instead", d, r.GetDirection())
	}
}

func Test_Move(t *testing.T) {
	id, x, y, d := 0, 1, 1, travel.East
	r := New(id, x, y, d)
	move := -3
	r.Move(move, move, travel.North)
	if r.positionX != x+move || r.positionY != y+move || r.direction != travel.North {
		t.Fatalf("expected Move to have moved robot to %d %d %s - moved instead to %v", x+move, y+move, travel.North, r)
	}
}

func Test_String(t *testing.T) {
	id, x, y, d := 0, 1, 1, travel.East
	r := New(id, x, y, d)
	expected := fmt.Sprintf("%d %d %s", x, y, d)
	if expected != r.String() {
		t.Fatalf("expected String to have produced %s - instead got %s", expected, r.String())
	}
}
