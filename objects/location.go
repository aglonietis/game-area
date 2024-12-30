package objects

import (
	"math"
	"math/rand"
	"soul/utilities"
)

type Location struct {
	World *World
	X     int
	Y     int
}

func NewLocation(world *World) *Location {
	return &Location{
		X:     0,
		Y:     0,
		World: world,
	}
}

func (l *Location) findNearbySouls() []*Soul {
	return l.World.FindNearbySouls(l)
}

// IsNearby Checks if location is less than 5 distances away
func (l *Location) IsNearby(otherL *Location) bool {
	return math.Abs(float64(l.X-otherL.X)) < 1 ||
		math.Abs(float64(l.Y-otherL.Y)) < 1
}

func (l *Location) MoveRandomly() {
	l.X += l.CalculateMovement(l.X < (l.World.Radius*2-1), l.X > 0)
	l.Y += l.CalculateMovement(l.Y < (l.World.Radius*2-1), l.Y > 0)
}

func (l *Location) CalculateMovement(canMovePositive, canMoveNegative bool) int {
	if canMovePositive && canMoveNegative {
		return 2*rand.Intn(2) - 1 // Randomly -1 or 1
	}
	return rand.Intn(2) * (2*utilities.BoolToInt(canMovePositive) - 1)
}
