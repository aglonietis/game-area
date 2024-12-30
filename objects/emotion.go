package objects

type Emotion struct {
	Min     int
	Max     int
	Current int
	Name    string
}

func NewEmotion(name string) Emotion {
	return Emotion{
		Name:    name,
		Min:     -100,
		Max:     100,
		Current: 0,
	}
}

func (e Emotion) Increase(amount int) {
	e.Current += amount
}

func (e Emotion) Decrease(amount int) {
	e.Current -= amount
}
