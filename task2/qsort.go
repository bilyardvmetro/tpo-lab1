package task2

import "fmt"

type Event struct {
	Name  string
	L, R  int
	I, J  int
	Pivot int
	Val   int
}

func (e Event) String() string {
	switch e.Name {
	case "START", "DONE":
		return e.Name
	case "CALL":
		return fmt.Sprintf("CALL(l=%d,r=%d)", e.L, e.R)
	case "P":
		return fmt.Sprintf("P(pivot=%d)", e.Pivot)
	case "C":
		return fmt.Sprintf("C(j=%d,pivot=%d)", e.J, e.Pivot)
	case "SW":
		return fmt.Sprintf("SW(i=%d,j=%d)", e.I, e.J)
	case "PS":
		return fmt.Sprintf("PS(pos=%d)", e.I)
	default:
		return e.Name
	}
}

type Tracer func(Event)

func emit(t Tracer, e Event) {
	if t != nil {
		t(e)
	}
}

func QuickSortInts(a []int, t Tracer) {
	emit(t, Event{Name: "START"})
	if len(a) > 1 {
		quickSort(a, 0, len(a)-1, t)
	}
	emit(t, Event{Name: "DONE"})
}

func quickSort(a []int, l, r int, t Tracer) {
	emit(t, Event{Name: "CALL", L: l, R: r})

	if l >= r {
		return
	}

	pivot := a[r]
	emit(t, Event{Name: "P", Pivot: pivot})

	i := l
	for j := l; j < r; j++ {
		emit(t, Event{Name: "C", J: j, Pivot: pivot})
		if a[j] < pivot {
			emit(t, Event{Name: "SW", I: i, J: j})
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	emit(t, Event{Name: "SW", I: i, J: r})
	a[i], a[r] = a[r], a[i]

	emit(t, Event{Name: "PS", I: i})

	quickSort(a, l, i-1, t)
	quickSort(a, i+1, r, t)
}
