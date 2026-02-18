package task3

import (
	"fmt"
	"io"
	"os"
)

type Question string

type Entity struct {
	Name  string
	State State
	Race  Race

	Out io.Writer
}

func NewEntity(name string, state State, race Race) *Entity {
	return &Entity{Name: name, State: state, Race: race, Out: os.Stdout}
}

func (e *Entity) SetOut(w io.Writer) {
	e.Out = w
}

type HumanEntity struct{ E *Entity }

func (h *HumanEntity) Entity() *Entity {
	return h.E
}

func (h *HumanEntity) BeatenBy(c Creature) {
	fmt.Fprintln(h.E.Out, "Человек", h.E.Name, "ударен существом", c.Entity().Name)
}

type CreatureEntity struct{ E *Entity }

func (c *CreatureEntity) Entity() *Entity {
	return c.E
}

func (c *CreatureEntity) Beat(h Human) {
	fmt.Fprintln(c.E.Out, "Существо", c.E.Name, "ударил человека", h.Entity().Name)
}

func (c *CreatureEntity) Run() {
	c.E.State = Running
	fmt.Fprintln(c.E.Out, "Существо", c.E.Name, "бежит")
}

func (c *CreatureEntity) Sit() {
	c.E.State = Sitting
	fmt.Fprintln(c.E.Out, "Существо", c.E.Name, "сел")
}

func (c *CreatureEntity) GetTired() {
	c.E.State = Tired
	fmt.Fprintln(c.E.Out, "Существо", c.E.Name, "устал")
}

func (c *CreatureEntity) GetDistracted() {
	c.E.State = Distracted
	fmt.Fprintln(c.E.Out, "Существо", c.E.Name, "отвлекся")
}

func (c *CreatureEntity) SolveQuestions(q []Question) {
	c.E.State = SolvingQuestions
	for _, question := range q {
		fmt.Fprintln(c.E.Out, "Существо", c.E.Name, "решает вопрос:", question)
	}
}

type Game struct {
	h Human
	c Creature
}

func InitGame(h Human, c Creature) *Game {
	return &Game{h: h, c: c}
}

func (g *Game) Execute() {
	g.c.Beat(g.h)
	g.h.BeatenBy(g.c)
	g.c.Run()
}
