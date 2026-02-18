package task3

type Human interface {
	Entity() *Entity
	BeatenBy(c Creature)
}

type Creature interface {
	Entity() *Entity
	Beat(h Human)
	Run()
	Sit()
	GetTired()
	GetDistracted()
	SolveQuestions(q []Question)
}
