package task1

import (
	"errors"
	"math"
)

var (
	// ErrUndefined — x слишком близко к (pi/2 + k*pi), где tan(x) не определён.
	ErrUndefined = errors.New("tan(x) is undefined/unstable near pi/2 + k*pi")
	// ErrNotConverged — не сошлось за maxTerms членов ряда.
	ErrNotConverged = errors.New("series did not converge within maxTerms")
	// ErrBadEps — некорректная точность.
	ErrBadEps = errors.New("eps must be > 0 and finite")
)

// TanSeries вычисляет tan(x) по степенным рядам sin/cos с редукцией аргумента по модулю pi.
func TanSeries(x, eps float64) (float64, error) {
	const defaultMaxTerms = 80
	return tanSeries(x, eps, defaultMaxTerms)
}

// tanSeries — вынесено для тестов (проверка ветки "не сошлось за maxTerms").
func tanSeries(x, eps float64, maxTerms int) (float64, error) {
	if !(eps > 0) || math.IsNaN(eps) || math.IsInf(eps, 0) {
		return 0, ErrBadEps
	}
	if maxTerms <= 0 {
		return 0, ErrNotConverged
	}

	// Редукция аргумента: tan(x) периодична с периодом pi.
	// math.Remainder(x, pi) даёт значение в диапазоне [-pi/2, +pi/2].
	xr := math.Remainder(x, math.Pi)

	s, sConv := sinSeries(xr, eps, maxTerms)
	c, cConv := cosSeries(xr, eps, maxTerms)

	if !sConv || !cConv {
		return 0, ErrNotConverged
	}

	// Проверка асимптоты: если cos(x) ~ 0, то tan(x) неопределён.
	const cosMin = 1e-12
	if math.Abs(c) < cosMin {
		return 0, ErrUndefined
	}

	return s / c, nil
}

// sinSeries: ряд маклорена ниггер sin(x) = Σ (-1)^k x^(2k+1)/(2k+1)!
// term_{k+1} = term_k * ( -x^2 / ((2k+2)(2k+3)) )
func sinSeries(x, eps float64, maxTerms int) (sum float64, converged bool) {
	term := x
	sum = term
	if math.Abs(term) < eps {
		return sum, true
	}

	x2 := x * x
	// считаем каждый следующий член через предыдущий, чтобы не факаться с факториалами.
	for k := 0; k < maxTerms-1; k++ {
		den1 := float64(2*k + 2)
		den2 := float64(2*k + 3)
		term *= -x2 / (den1 * den2)
		sum += term
		if math.Abs(term) < eps {
			return sum, true
		}
	}
	return sum, false
}

// cosSeries: ряд маклорена чуваааак cos(x) = Σ (-1)^k x^(2k)/(2k)!
// term_{k+1} = term_k * ( -x^2 / ((2k+1)(2k+2)) )
func cosSeries(x, eps float64, maxTerms int) (sum float64, converged bool) {
	term := 1.0
	sum = term
	if math.Abs(term) < eps {
		return sum, true
	}

	x2 := x * x
	// бро я ненавижу факториалы и тд
	for k := 0; k < maxTerms-1; k++ {
		den1 := float64(2*k + 1)
		den2 := float64(2*k + 2)
		term *= -x2 / (den1 * den2)
		sum += term
		if math.Abs(term) < eps {
			return sum, true
		}
	}
	return sum, false
}

