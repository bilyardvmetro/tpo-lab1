package task3

import (
	"bytes"
	"strings"
	"testing"
)

func mustContain(t *testing.T, out string, parts ...string) {
	t.Helper()
	for _, p := range parts {
		if !strings.Contains(out, p) {
			t.Fatalf("output must contain %q, got: %q", p, out)
		}
	}
}

func newEntityWithBuf(name string, st State, r Race) (*Entity, *bytes.Buffer) {
	var buf bytes.Buffer

	e := NewEntity(name, st, r)
	e.SetOut(&buf)

	return e, &buf
}

func TestHuman_EntityReturnsSamePointer(t *testing.T) {
	// t.Parallel()

	e, _ := newEntityWithBuf("Сергей Бессмертный", Sitting, HumanRace)
	h := &HumanEntity{E: e}

	if h.Entity() != e {
		t.Fatalf("HumanEntity.Entity() must return the same pointer")
	}
}

func TestCreature_EntityReturnsSamePointer(t *testing.T) {
	// t.Parallel()

	e, _ := newEntityWithBuf("Существо", Sitting, CreatureRace)
	c := &CreatureEntity{E: e}

	if c.Entity() != e {
		t.Fatalf("CreatureEntity.Entity() must return the same pointer")
	}
}

func TestCreature_Run_ChangesStateAndPrints(t *testing.T) {
	// t.Parallel()

	e, buf := newEntityWithBuf("Существо", Sitting, CreatureRace)
	c := &CreatureEntity{E: e}

	c.Run()

	if e.State != Running {
		t.Fatalf("state=%v want=%v", e.State, Running)
	}

	out := buf.String()
	mustContain(t, out, "Существо", e.Name, "бежит")
}

func TestCreature_Sit_ChangesStateAndPrints(t *testing.T) {
	// t.Parallel()

	e, buf := newEntityWithBuf("Существо", Running, CreatureRace)
	c := &CreatureEntity{E: e}

	c.Sit()

	if e.State != Sitting {
		t.Fatalf("state=%v want=%v", e.State, Sitting)
	}

	out := buf.String()
	mustContain(t, out, "Существо", e.Name, "сел")
}

func TestCreature_GetTired_ChangesStateAndPrints(t *testing.T) {
	// t.Parallel()

	e, buf := newEntityWithBuf("Существо", Sitting, CreatureRace)
	c := &CreatureEntity{E: e}

	c.GetTired()

	if e.State != Tired {
		t.Fatalf("state=%v want=%v", e.State, Tired)
	}

	out := buf.String()
	mustContain(t, out, "Существо", e.Name, "устал")
}

func TestCreature_GetDistracted_ChangesStateAndPrints(t *testing.T) {
	// t.Parallel()

	e, buf := newEntityWithBuf("Существо", Sitting, CreatureRace)
	c := &CreatureEntity{E: e}

	c.GetDistracted()

	if e.State != Distracted {
		t.Fatalf("state=%v want=%v", e.State, Distracted)
	}

	out := buf.String()
	mustContain(t, out, "Существо", e.Name, "отвлекся")
}

func TestCreature_SolveQuestions_PrintsEachQuestion(t *testing.T) {
	// t.Parallel()

	e, buf := newEntityWithBuf("Существо", Sitting, CreatureRace)
	c := &CreatureEntity{E: e}

	qs := []Question{"смысл жизни", "ультра-крикет"}
	c.SolveQuestions(qs)

	out := buf.String()
	mustContain(t, out, "Существо", e.Name, "решает вопрос:", string(qs[0]), string(qs[1]))

	if e.State != SolvingQuestions {
	    t.Fatalf("state=%v want=%v", e.State, SolvingQuestions)
	}
}

func TestCreature_Beat_Prints(t *testing.T) {
	// t.Parallel()

	creE, creBuf := newEntityWithBuf("Существо", Sitting, CreatureRace)
	humE, _ := newEntityWithBuf("Сергей Бессмертный", Sitting, HumanRace)

	c := &CreatureEntity{E: creE}
	h := &HumanEntity{E: humE}

	c.Beat(h)

	out := creBuf.String()
	mustContain(t, out, "Существо", creE.Name, "ударил", "человека", humE.Name)
}

func TestHuman_BeatenBy_Prints(t *testing.T) {
	// t.Parallel()

	humE, humBuf := newEntityWithBuf("Вася", Sitting, HumanRace)
	creE, _ := newEntityWithBuf("Существо", Sitting, CreatureRace)

	h := &HumanEntity{E: humE}
	c := &CreatureEntity{E: creE}

	h.BeatenBy(c)

	out := humBuf.String()
	mustContain(t, out, "Человек", humE.Name, "ударен", "существом", creE.Name)
}

func TestGame_Execute_PrintsAndMakesCreatureRun(t *testing.T) {
	// t.Parallel()

	creE, creBuf := newEntityWithBuf("Существо", Sitting, CreatureRace)
	humE, humBuf := newEntityWithBuf("Сергей Бессмертный", Sitting, HumanRace)

	c := &CreatureEntity{E: creE}
	h := &HumanEntity{E: humE}

	g := InitGame(h, c)
	g.Execute()

	if creE.State != Running {
		t.Fatalf("after Execute creature state=%v want=%v", creE.State, Running)
	}

	outC := creBuf.String()
	mustContain(t, outC, "Существо", creE.Name, "ударил", "бежит")

	outH := humBuf.String()
	mustContain(t, outH, "Человек", humE.Name, "ударен", "существом", creE.Name)
}