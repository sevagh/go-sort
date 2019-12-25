package bigO

import (
	"fmt"
	"math"
)

type complexity int

const (
	_       = iota
	O1      = iota
	ON      = iota
	ON2     = iota
	ON3     = iota
	OLogN   = iota
	ONLogN  = iota
	OLambda = iota
)

var fitCurves = []complexity{O1, ON, ON2, ON3, OLogN, ONLogN}

type BigO struct {
	coef float64
	rms  float64
	cplx complexity
}

func (b *BigO) String() string {
	return complexity2Str(b.cplx)
}

func complexity2Str(cplx complexity) string {
	switch cplx {
	case ON:
		return "O(N)"
	case ON2:
		return "O(N^2)"
	case ON3:
		return "O(N^3)"
	case OLogN:
		return "O(lg(N))"
	case ONLogN:
		return "O(Nlg(N))"
	case O1:
		return "O(1)"
	default:
		return "f(n)"
	}
}

func fittingCurve(cplx complexity) func(float64) float64 {
	switch cplx {
	case ON:
		return func(n float64) float64 {
			return n
		}
	case ON2:
		return func(n float64) float64 {
			return n * n
		}
	case ON3:
		return func(n float64) float64 {
			return n * n * n
		}
	case OLogN:
		return func(n float64) float64 {
			return math.Log2(n)
		}
	case ONLogN:
		return func(n float64) float64 {
			return n * math.Log2(n)
		}
	default:
		return func(n float64) float64 {
			return 1.0
		}
	}
}

func minimalLeastSq(n []int, times []float64, fittingCurve func(float64) float64) BigO {
	sigmaGn := 0.0
	sigmaGnSquared := 0.0
	sigmaTime := 0.0
	sigmaTimeGn := 0.0

	floatN := float64(len(n))

	// Calculate least square fitting parameter
	for i := 0; i < len(n); i++ {
		gnI := fittingCurve(float64(n[i]))
		sigmaGn += gnI
		sigmaGnSquared += gnI * gnI
		sigmaTime += times[i]
		sigmaTimeGn += times[i] * gnI
	}

	var result BigO
	result.cplx = OLambda

	// Calculate complexity.
	result.coef = sigmaTimeGn / sigmaGnSquared

	// Calculate RMS
	rms := 0.0
	for i := 0; i < len(n); i++ {
		fit := result.coef * fittingCurve(float64(n[i]))
		rms += math.Pow(times[i]-fit, 2)
	}

	// Normalized RMS by the mean of the observed values
	mean := sigmaTime / floatN
	result.rms = math.Sqrt(rms/floatN) / mean

	return result
}

func Estimate(n []int, times []float64) (BigO, error) {
	var bestFit BigO

	if len(n) != len(times) {
		return bestFit, fmt.Errorf("length mismatch between n and times slices")
	}
	if len(n) < 2 {
		return bestFit, fmt.Errorf("need at least 2 runs")
	}

	bestFit = minimalLeastSq(n, times, fittingCurve(O1))
	bestFit.cplx = O1

	for _, fit := range fitCurves {
		currentFit := minimalLeastSq(n, times, fittingCurve(fit))
		if currentFit.rms < bestFit.rms {
			bestFit = currentFit
			bestFit.cplx = fit
		}
	}

	return bestFit, nil
}
