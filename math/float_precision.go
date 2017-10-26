package math

import "math"

func round(n float64) int {
    return int(n + math.Copysign(0.5, n))
}

func TruncateFloat(f float64, precision int) float64 {
	p := math.Pow(10, float64(precision))
	return float64(round(f*p)) / p
}
