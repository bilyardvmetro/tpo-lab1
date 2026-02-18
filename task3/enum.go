package task3

type State string

const (
	Running          State = "бежит"
	Sitting          State = "сидит"
	Tired            State = "устал"
	SolvingQuestions State = "решает вопросы"
	Distracted       State = "отвлечен"
)

type Race string

const (
	HumanRace    Race = "человек"
	CreatureRace Race = "существо"
)
