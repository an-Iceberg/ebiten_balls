package utils

import "math/rand"

func Clamp(number, min, max float64) float64 {
	if number < min {
		return min
	}
	if number > max {
		return max
	}
	return number
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
