package robot

import (
	"fmt"

	"github.com/juubisnake/mars-rover/internal/pkg/travel"
)

// Robot is a representation of a robot that is capable of moving across a surface.
type Robot struct {
	id        int
	positionX int
	positionY int
	direction travel.Direction
}

// New creates an instance of a robot that holds information of where it is positioned
// within a surface.
func New(id, x, y int, direction travel.Direction) *Robot {
	return &Robot{
		id:        id,
		positionX: x,
		positionY: y,
		direction: direction,
	}
}

// GetID returns the ID of the given robot.
func (r *Robot) GetID() int {
	return r.id
}

// GetX returns the x coordinate of the given robot.
func (r *Robot) GetX() int {
	return r.positionX
}

// GetY returns the y coordinate of the given robot.
func (r *Robot) GetY() int {
	return r.positionY
}

// GetDirection returns the direction the given robot is facing.
func (r *Robot) GetDirection() travel.Direction {
	return r.direction
}

// Move translates the robot via the given coordinates and direction and updates
// its current position.
// I.E Move(1, 2, direction.East) will move the robot by 1 on the x-axis, 2 on the y-axis
// and point the robot east.
func (r *Robot) Move(x, y int, direction travel.Direction) {
	r.positionX += x
	r.positionY += y
	r.direction = direction
}

// String is a representation of a given robot in the form of X Y DIRECTION.
// I.E A robot that has x=2 y=4 direction=west will output 2 4 W.
func (r *Robot) String() string {
	return fmt.Sprintf("%d %d %s", r.positionX, r.positionY, r.direction)
}
