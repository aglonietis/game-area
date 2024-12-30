package objects

import (
	"math/rand"
)

type Soul struct {
	ID              int
	Name            string
	Initials        string
	Emotions        *Emotions
	ShortTermMemory Memory
	LongTermMemory  Memory
	Location        *Location
}

func NewSoul(name string, world *World) *Soul {
	return &Soul{
		ID:              rand.Int(),
		Name:            name,
		Initials:        string(name[0]),
		Emotions:        NewEmotions(),
		ShortTermMemory: NewMemory(),
		LongTermMemory:  NewMemory(),
		Location:        NewLocation(world),
	}
}

func (s *Soul) Act() {
	nearbySoulsWithSelf := s.Location.findNearbySouls()

	var nearbySouls []*Soul
	for _, soul := range nearbySoulsWithSelf {
		if soul != s {
			nearbySouls = append(nearbySouls, soul)
		}
	}

	if len(nearbySouls) > 0 {
		unluckySoul := nearbySouls[rand.Intn(len(nearbySouls))]
		s.Hit(unluckySoul)
	} else {
		s.Travel()
	}
}

func (s *Soul) MoveRandomly() {
	s.Location.MoveRandomly()
}

func (s *Soul) MoveAway() {
	s.MoveRandomly()
}

func (s *Soul) Travel() {
	s.Emotions.Loneliness.Increase(1)
	s.Emotions.Joy.Decrease(1)
	s.MoveRandomly()
}

func (s *Soul) Hit(otherS *Soul) {
	s.Emotions.Joy.Increase(1)
	s.Emotions.Loneliness.Decrease(1)
	otherS.GetHit()
}

func (s *Soul) GetHit() {
	s.Emotions.Joy.Decrease(1)
	s.MoveAway()
}
