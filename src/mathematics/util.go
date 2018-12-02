package mathematics

import (
	"math"
)

func round(value float64, roundOn float64, places int) float64 {
	pow := math.Pow(10, float64(places))
	digit := pow * value
	_, div := math.Modf(digit)
	div = math.Copysign(div, value)
	roundOn = math.Copysign(roundOn, value)

	var round float64
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}

func findElementFrequency(X []int) map[int]int {
	m := make(map[int]int)
	for _, x := range X {
		m[x] = 0
	}
	for _, x := range X {
		for k := range m {
			if x == k {
				m[k]++
			}
		}
	}
	return m
}

func findCommonIntegers(X, Y []int) []int {
	var common []int

	for i, x := range X {
		for _, y := range Y {
			if x == y {
				common = append(common, x)
				X = append(X[:i], 0)
				break
			}
		}
	}
	return common
}
