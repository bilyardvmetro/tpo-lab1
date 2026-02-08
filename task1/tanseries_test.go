package task1

import (
	"math"
	"testing"
)

func almostEqualAbsRel(got, want, absTol, relTol float64) bool {
	diff := math.Abs(got - want)
	if diff <= absTol {
		return true
	}
	// погрешность
	scale := math.Max(1.0, math.Abs(want))
	return diff/scale <= relTol
}

func TestTanSeries_BasicPoints(t *testing.T) {
	// t.Parallel()

	// Точки вдали от асимптот (pi/2 + k*pi)
	tests := map[float64]float64{
		0: 0,
		1e-12: 1e-12,
		1e-6: 1e-6,
		-1e-6: -1e-6,
		0.1: 0.100_334_672_085,
		-0.1: -0.100_334_672_085,
		0.5: 0.546_302_489_843,
		-0.5: -0.546_302_489_843,
		1.0: 1.557_407_724_654,
		-1.0: -1.557_407_724_654,
	}

	const eps = 1e-15
	for x, want := range tests {
		got, err := TanSeries(x, eps)
		if err != nil {
			t.Fatalf("x=%v: unexpected err: %v", x, err)
		}

		if !almostEqualAbsRel(got, want, 1e-12, 1e-12) {
			t.Fatalf("x=%v: got=%0.17g want=%0.17g", x, got, want)
		}
	}
}

func TestTanSeries_RangeReductionPeriodPi(t *testing.T) {
	// t.Parallel()

	// tan(x) == tan(x + k*pi) через редукцию аргумента.
	x := 0.3
	k := 123456.0
	x2 := x + k*math.Pi

	got1, err := TanSeries(x, 1e-15)
	if err != nil {
		t.Fatalf("unexpected err for x: %v", err)
	}
	got2, err := TanSeries(x2, 1e-15)
	if err != nil {
		t.Fatalf("unexpected err for x2: %v", err)
	}

	if !almostEqualAbsRel(got1, got2, 1e-10, 1e-10) {
		t.Fatalf("periodicity: got1=%0.17g got2=%0.17g", got1, got2)
	}
}

func TestTanSeries_NearAsymptoteReturnsErrUndefined(t *testing.T) {
	// t.Parallel()

	// Очень близко к pi/2: cos(x) ~ 0
	x := math.Pi/2 - 1e-16

	_, err := TanSeries(x, 1e-18)
	if err == nil {
		t.Fatalf("expected error near asymptote, got nil")
	}
	if err != ErrUndefined {
		t.Fatalf("expected ErrUndefined, got: %v", err)
	}
}

func TestTanSeries_BadEps(t *testing.T) {
	// t.Parallel()

	_, err := TanSeries(0.1, 0)
	if err != ErrBadEps {
		t.Fatalf("expected ErrBadEps for eps=0, got: %v", err)
	}

	_, err = TanSeries(0.1, math.NaN())
	if err != ErrBadEps {
		t.Fatalf("expected ErrBadEps for eps=NaN, got: %v", err)
	}

	_, err = TanSeries(0.1, math.Inf(1))
	if err != ErrBadEps {
		t.Fatalf("expected ErrBadEps for eps=+Inf, got: %v", err)
	}
}

func TestTanSeries_NotConvergedWhenMaxTermsTooSmall(t *testing.T) {
	// t.Parallel()

	// очень маленький eps и очень маленький maxTerms.
	// x умеренный, чтобы требовалось больше членов.
	x := 1.2

	_, err := tanSeries(x, 1e-30, 2)
	if err == nil {
		t.Fatalf("expected ErrNotConverged, got nil")
	}
	if err != ErrNotConverged {
		t.Fatalf("expected ErrNotConverged, got: %v", err)
	}
}

func TestSinCosSeries_EarlyStop(t *testing.T) {
	// t.Parallel()

	// мгонвенное вычисление при очень малом x:
	// term уже < eps после 1-2 шагов.
	x := 1e-20
	eps := 1e-18

	s, sConv := sinSeries(x, eps, 50)
	c, cConv := cosSeries(x, eps, 50)

	if !sConv || !cConv {
		t.Fatalf("expected early convergence, got sConv=%v cConv=%v", sConv, cConv)
	}

	// sin(x) ~ x, cos(x) ~ 1
	if math.Abs(s-x) > 1e-20 {
		t.Fatalf("sin approx: got=%g want~%g", s, x)
	}
	if math.Abs(c-1.0) > 1e-18 {
		t.Fatalf("cos approx: got=%g want~1", c)
	}
}
