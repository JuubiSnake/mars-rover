package plateau

import "testing"

func Test_New_X_InvalidUpperBound(t *testing.T) {
	_, err := New(-1, 0)
	if err == nil {
		t.Fatalf("expected New to fail with x coordinate that is less than lower bound %d", lowerBoundX)
	}
	ue, ok := err.(*UpperBoundsError)
	if !ok {
		t.Fatalf("expected err to be UpperBoundsError - got %T instead", err)
	}
	if ue.Coordinate != "x" || ue.Value != -1 || ue.LowerBound != lowerBoundX {
		t.Fatalf("expected error to fail on x coordinate of -1 which is less then lower bound for x %d - instead got: %v", lowerBoundX, err)
	}
}

func Test_New_Y_InvalidUpperBound(t *testing.T) {
	_, err := New(0, -1)
	if err == nil {
		t.Fatalf("expected New to fail with y coordinate that is less than lower bound %d", lowerBoundY)
	}
	ue, ok := err.(*UpperBoundsError)
	if !ok {
		t.Fatalf("expected err to be UpperBoundsError - got %T instead", err)
	}
	if ue.Coordinate != "y" || ue.Value != -1 || ue.LowerBound != lowerBoundY {
		t.Fatalf("expected error to fail on y coordinate of -1 which is less then lower bound for y %d - instead got: %v", lowerBoundY, err)
	}
}

func Test_New_Bound(t *testing.T) {
	x := 0
	y := 0
	s, err := New(x, y)
	if err != nil {
		t.Fatalf("New should have been valid with upper-bound coordinates of 0,0 - instead got the following error: %v", err)
	}
	if s.UpperBoundX != x || s.UpperBoundY != y || s.LowerBoundX != lowerBoundX || s.LowerBoundY != lowerBoundY {
		t.Fatalf("New should have produced a surface that was lower bounded to %d,%d and upper bounded to %d,%d - instead it produced: %v", lowerBoundX, lowerBoundY, s.UpperBoundX, s.UpperBoundY, s)
	}
}

func Test_IsOutOfBounds(t *testing.T) {
	x := 0
	y := 0
	s, err := New(x, y)
	if err != nil {
		t.Fatalf("New should have been valid with upper-bound coordinates of 0,0 - instead got the following error: %v", err)
	}

	if !s.IsOutOfBounds(x+1, y) {
		t.Fatalf("x=%d should have been out-of-bounds", x+1)
	}
	if !s.IsOutOfBounds(x+1, y) {
		t.Fatalf("x=%d should have been out-of-bounds", x-1)
	}
	if !s.IsOutOfBounds(x, y+1) {
		t.Fatalf("y=%d should have been out-of-bounds", y+1)
	}
	if !s.IsOutOfBounds(x, y-1) {
		t.Fatalf("y=%d should have been out-of-bounds", y-1)
	}
	if s.IsOutOfBounds(x, y) {
		t.Fatalf("x=%d y=%d should not have been out-of-bounds", x, y)
	}
}
