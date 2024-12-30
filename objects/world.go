package objects

type World struct {
	Souls  []*Soul
	Radius int
}

func NewWorld(size int) *World {
	return &World{
		Radius: size,
	}
}

func (w *World) AddSoul(soul *Soul) {
	w.Souls = append(w.Souls, soul)
}

func (w *World) FindNearbySouls(l *Location) []*Soul {
	var nearbySouls []*Soul
	for _, soul := range w.Souls {
		if soul.Location.IsNearby(l) {
			nearbySouls = append(nearbySouls, soul)
		}
	}

	return nearbySouls
}
