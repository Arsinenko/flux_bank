package behavior

import "math"

func Linear(value, min, max float64) float64 {
	if value < min {
		return 0
	}
	if value > max {
		return 1
	}
	return (value - min) / (max - min)
}

func ResponseCurve(value float64, exponent float64) float64 {
	return math.Pow(math.Max(0, math.Min(1, value)), exponent)
}
