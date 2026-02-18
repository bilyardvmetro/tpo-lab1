package task2

import (
	"reflect"
	"strings"
	"testing"
)

func recordTrace() (Tracer, *[]string) {
	var trace []string
	return func(e Event) {
		trace = append(trace, e.String())
	}, &trace
}

func TestQuickSort_Trace_312(t *testing.T) {
	// t.Parallel()

	a := []int{3, 1, 2}
	tr, got := recordTrace()

	QuickSortInts(a, tr)

	wantTrace := []string{
		"START",
		"CALL(l=0,r=2)",
		"P(pivot=2)",
		"C(j=0,pivot=2)",
		"C(j=1,pivot=2)",
		"SW(i=0,j=1)",
		"SW(i=1,j=2)",
		"PS(pos=1)",

		"CALL(l=0,r=0)",
		"CALL(l=2,r=2)",
		"DONE",
	}

	if !reflect.DeepEqual(*got, wantTrace) {
		t.Fatalf("trace mismatch\n got:  %#v\n want: %#v", *got, wantTrace)
	}

	if !reflect.DeepEqual(a, []int{1, 2, 3}) {
		t.Fatalf("sorted array mismatch: got=%v", a)
	}
}

func TestQuickSort_AlreadySorted(t *testing.T) {
	// t.Parallel()

	a := []int{1, 2, 3}
	QuickSortInts(a, nil)

	if !reflect.DeepEqual(a, []int{1, 2, 3}) {
		t.Fatalf("sorted array mismatch: got=%v", a)
	}
}

func TestQuickSort_Reverse(t *testing.T) {
	// t.Parallel()

	a := []int{3, 2, 1}
	QuickSortInts(a, nil)

	if !reflect.DeepEqual(a, []int{1, 2, 3}) {
		t.Fatalf("sorted array mismatch: got=%v", a)
	}
}

func TestQuickSort_Duplicates(t *testing.T) {
	// t.Parallel()

	a := []int{2, 1, 2, 1}
	QuickSortInts(a, nil)

	if !reflect.DeepEqual(a, []int{1, 1, 2, 2}) {
		t.Fatalf("sorted array mismatch: got=%v", a)
	}
}

func TestQuickSort_Trace_Empty(t *testing.T) {
	// t.Parallel()

	a := []int{}
	tr, got := recordTrace()

	QuickSortInts(a, tr)

	want := []string{
		"START",
		"DONE",
	}

	if !reflect.DeepEqual(*got, want) {
		t.Fatalf("trace mismatch\n got: %v\n want: %v", *got, want)
	}
}

func TestQuickSort_Single(t *testing.T) {
	// t.Parallel()

	a := []int{42}
	tr, got := recordTrace()

	QuickSortInts(a, tr)

	want := []string{
		"START",
		"DONE",
	}

	if !reflect.DeepEqual(*got, want) {
		t.Fatalf("trace mismatch\n got: %v\n want: %v", *got, want)
	}
}

func TestQuickSort_Trace_AllEqual(t *testing.T) {
	t.Parallel()

	a := []int{7, 7, 7}
	tr, got := recordTrace()

	QuickSortInts(a, tr)

	want := []string{
		"START",
		"CALL(l=0,r=2)",
		"P(pivot=7)",
		"C(j=0,pivot=7)",
		"C(j=1,pivot=7)",
		"SW(i=0,j=2)",
		"PS(pos=0)",

		"CALL(l=0,r=-1)",
		"CALL(l=1,r=2)",
		"P(pivot=7)",
		"C(j=1,pivot=7)",
		"SW(i=1,j=2)",
		"PS(pos=1)",
		"CALL(l=1,r=0)",
		"CALL(l=2,r=2)",

		"DONE",
	}

	if !reflect.DeepEqual(*got, want) {
		t.Fatalf("trace mismatch\n got: %v\n want: %v", *got, want)
	}
}

func TestQuickSort_Trace_Negatives(t *testing.T) {
	t.Parallel()

	a := []int{0, -1, 5, -10, 3}
	tr, got := recordTrace()

	QuickSortInts(a, tr)

	trace := *got
	has := func(prefix string) bool {
		for _, s := range trace {
			if strings.HasPrefix(s, prefix) {
				return true
			}
		}
		return false
	}

	if trace[0] != "START" {
		t.Fatalf("trace must start with START")
	}
	if trace[len(trace)-1] != "DONE" {
		t.Fatalf("trace must end with DONE")
	}

	if !has("P(") {
		t.Fatalf("expected pivot selection in trace")
	}
	if !has("C(") {
		t.Fatalf("expected comparisons in trace")
	}
	if !has("PS(") {
		t.Fatalf("expected pivot placement in trace")
	}
}
