package objects

type Emotions struct {
	Joy         Emotion
	Sadness     Emotion
	Fear        Emotion
	Anger       Emotion
	Disgust     Emotion
	Surprise    Emotion
	Love        Emotion
	Shame       Emotion
	Guilt       Emotion
	Pride       Emotion
	Envy        Emotion
	Jealousy    Emotion
	Hope        Emotion
	Gratitude   Emotion
	Compassion  Emotion
	Contempt    Emotion
	Frustration Emotion
	Anxiety     Emotion
	Relief      Emotion
	Loneliness  Emotion
}

func NewEmotions() *Emotions {
	return &Emotions{
		Joy:         NewEmotion("Joy"),
		Sadness:     NewEmotion("Sadness"),
		Fear:        NewEmotion("Fear"),
		Anger:       NewEmotion("Anger"),
		Disgust:     NewEmotion("Disgust"),
		Surprise:    NewEmotion("Surprise"),
		Love:        NewEmotion("Love"),
		Shame:       NewEmotion("Shame"),
		Guilt:       NewEmotion("Guilt"),
		Pride:       NewEmotion("Pride"),
		Envy:        NewEmotion("Envy"),
		Jealousy:    NewEmotion("Jealousy"),
		Hope:        NewEmotion("Hope"),
		Gratitude:   NewEmotion("Gratitude"),
		Compassion:  NewEmotion("Compassion"),
		Contempt:    NewEmotion("Contempt"),
		Frustration: NewEmotion("Frustration"),
		Anxiety:     NewEmotion("Anxiety"),
		Relief:      NewEmotion("Relief"),
		Loneliness:  NewEmotion("Loneliness"),
	}
}
