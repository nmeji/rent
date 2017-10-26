package math

func SumInt(n []int) int {
	sum := 0
	for _, i := range n {
		sum = sum + i
	}
	return sum
}

func SumFloat(n []float64) float64 {
	sum := float64(0)
	for _, i := range n {
		sum = sum + i
	}
	return sum
}

func Avg(x,y int) float64 {
	return float64(x)/float64(y)
}

