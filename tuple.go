package main

import "math"

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// Add a tuple to the instance and return the result. Note that the sum of two
// vectors is a vector, the sum of a point and a vector is a point, and the sum
// of two points is nonsensical.
func (t Tuple) Add(other Tuple) Tuple {
	return Tuple{
		t.X + other.X,
		t.Y + other.Y,
		t.Z + other.Z,
		t.W + other.W,
	}
}

// Find the cross product of this tuple and another.
func (t Tuple) Cross(other Tuple) Tuple {
	return MakeVector(
		t.Y*other.Z-t.Z*other.Y,
		t.Z*other.X-t.X*other.Z,
		t.X*other.Y-t.Y*other.X,
	)
}

// Divide the tuple by the provided factor and return the result. This is
// equivalent to dividing all the tuple's components by the factor.
func (t Tuple) Divide(scale float64) Tuple {
	return Tuple{t.X / scale, t.Y / scale, t.Z / scale, t.W / scale}
}

// Compute the dot product of two tuples.
func (t Tuple) Dot(other Tuple) float64 {
	return t.X*other.X +
		t.Y*other.Y +
		t.Z*other.Z +
		t.W*other.W
}

func (t Tuple) Equals(other Tuple) bool {
	return Float64Equal(t.X, other.X) &&
		Float64Equal(t.Y, other.Y) &&
		Float64Equal(t.Z, other.Z) &&
		Float64Equal(t.W, other.W)
}

func (t Tuple) IsPoint() bool {
	return t.W == 1
}

func (t Tuple) IsVector() bool {
	return t.W == 0
}

// Get the total distance represented by a vector.
func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

// Multiply the tuple by the provided factor and return the result. This is
// equivalent to multiplying all the tuple's components by the factor.
func (t Tuple) Multiply(scale float64) Tuple {
	return Tuple{t.X * scale, t.Y * scale, t.Z * scale, t.W * scale}
}

// Negate a tuple and return the result. This is equivalent to subtracting the
// tuple from the zero vector.
func (t Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

// Get the normalized version of a tuple. The normalized tuple maintains the
// ratio of the tuple's components to each other but has a magnitude of 1.
func (t Tuple) Normalized() Tuple {
	magnitude := t.Magnitude()

	return Tuple{
		t.X / magnitude,
		t.Y / magnitude,
		t.Z / magnitude,
		t.W / magnitude,
	}
}

// Subtract a tuple from the instance and return the result. Note that
// subtracting a point from a point is a vector, subtracting a vector from a
// point is a point, subtracting a vector from a vector is a vector, and
// subtracting a point from a vector is nonsensical.
func (t Tuple) Subtract(other Tuple) Tuple {
	return Tuple{
		t.X - other.X,
		t.Y - other.Y,
		t.Z - other.Z,
		t.W - other.W,
	}
}

func MakePoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func MakeVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}
