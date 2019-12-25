package bigO

import (
	"testing"
)

func TestConstant(t *testing.T) {
	n := []int{1, 2, 3, 4}

	times := []float64{7.0, 6.0, 8.0, 7.0}
	fit, _ := Estimate(n, times)

	betterTimes := []float64{5.0, 5.0, 5.0, 5.0}
	betterFit, _ := Estimate(n, betterTimes)

	if fit.String() != "O(1)" || betterFit.String() != "O(1)" {
		t.Fatalf("wrong complexity for squared, got %s or %s", fit.String(), betterFit.String())
	}

	if betterFit.rms > fit.rms || betterFit.coef > fit.coef {
		t.Errorf("expected second fit to be better, %+v vs %+v", betterFit, fit)
	}
}

func TestLinear(t *testing.T) {
	n := []int{1, 2, 3, 4}

	times := []float64{4, 3, 9, 13}
	fit, _ := Estimate(n, times)

	betterTimes := []float64{3, 6, 9, 12}
	betterFit, _ := Estimate(n, betterTimes)

	if fit.String() != "O(N)" || betterFit.String() != "O(N)" {
		t.Fatalf("wrong complexity for squared, got %s or %s", fit.String(), betterFit.String())
	}
}

func TestSquared(t *testing.T) {
	n := []int{1, 2, 3, 4}

	times := []float64{1, 5, 7, 16}
	fit, _ := Estimate(n, times)

	betterTimes := []float64{1, 4, 9, 16}
	betterFit, _ := Estimate(n, betterTimes)

	if fit.String() != "O(N^2)" || betterFit.String() != "O(N^2)" {
		t.Fatalf("wrong complexity for squared, got %s or %s", fit.String(), betterFit.String())
	}
}
