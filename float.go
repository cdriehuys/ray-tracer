package main

import "math"

// Margin of error for floating point comparisons
const floatEpsilon = 1e-9

// Determine if two floats are approximately equal to each other.
func Float64Equal(x, y float64) bool {
	return math.Abs(x - y) < floatEpsilon
}
