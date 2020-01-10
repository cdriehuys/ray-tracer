package main

import (
	"math"
	"testing"
)

type tupleOperationTest struct {
	name  string
	base  Tuple
	other Tuple
	want  Tuple
}

func TestTuple_Add(t *testing.T) {
	tests := []tupleOperationTest{
		{
			"vector addition",
			MakeVector(3, -2, 5),
			MakeVector(-2, 3, 1),
			MakeVector(1, 1, 6),
		},
		{
			"vector point addition",
			MakeVector(3, -2, 5),
			MakePoint(-2, 3, 1),
			MakePoint(1, 1, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			if got := tt.base.Add(tt.other); got != tt.want {
				t1.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Cross(t *testing.T) {
	tests := []tupleOperationTest{
		{
			"A x B",
			MakeVector(1, 2, 3),
			MakeVector(2, 3, 4),
			MakeVector(-1, 2, -1),
		},
		{
			"B x A",
			MakeVector(2, 3, 4),
			MakeVector(1, 2, 3),
			MakeVector(1, -2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			if got := tt.base.Cross(tt.other); got != tt.want {
				t1.Errorf("Cross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Divide(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		scale float64
		want  Tuple
	}{
		{
			name:  "divide by scalar",
			tuple: Tuple{1, -2, 3, -4},
			scale: 2,
			want:  Tuple{0.5, -1, 1.5, -2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			if got := tt.tuple.Divide(tt.scale); got != tt.want {
				t1.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Dot(t *testing.T) {
	tests := []struct {
		name  string
		base  Tuple
		other Tuple
		want  float64
	}{
		{
			name:  "dot product",
			base:  MakeVector(1, 2, 3),
			other: MakeVector(2, 3, 4),
			want:  20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			if got := tt.base.Dot(tt.other); got != tt.want {
				t1.Errorf("Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Equals(t1 *testing.T) {
	tests := []struct {
		name  string
		base  Tuple
		other Tuple
		want  bool
	}{
		{
			"equal points",
			MakePoint(1.1, -2.2, 3.3),
			MakePoint(1.1, -2.2, 3.3),
			true,
		},
		{
			"equal vectors",
			MakeVector(-1.1, 2.2, -3.3),
			MakeVector(-1.1, 2.2, -3.3),
			true,
		},
		{
			"different x",
			MakePoint(1, 0, 0),
			MakePoint(0, 0, 0),
			false,
		},
		{
			"different y",
			MakePoint(0, 1, 0),
			MakePoint(0, 0, 0),
			false,
		},
		{
			"different z",
			MakePoint(0, 0, 1),
			MakePoint(0, 0, 0),
			false,
		},
		{
			"point and vector",
			MakePoint(1.1, -2.2, 3.3),
			MakeVector(1.1, -2.2, 3.3),
			false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.base.Equals(tt.other); got != tt.want {
				t1.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_IsPoint(t1 *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		want  bool
	}{
		{
			name:  "point",
			tuple: MakePoint(4.3, -4.2, 3.1),
			want:  true,
		},
		{
			name:  "vector",
			tuple: MakeVector(4.3, -4.2, 3.1),
			want:  false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.tuple.IsPoint(); got != tt.want {
				t1.Errorf("IsPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_IsVector(t1 *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		want  bool
	}{
		{
			name:  "point",
			tuple: MakePoint(4.3, -4.2, 3.1),
			want:  false,
		},
		{
			name:  "vector",
			tuple: MakeVector(4.3, -4.2, 3.1),
			want:  true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.tuple.IsVector(); got != tt.want {
				t1.Errorf("IsVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Magnitude(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		want  float64
	}{
		{
			name:  "magnitude of x unit vector",
			tuple: MakeVector(1, 0, 0),
			want:  1,
		},
		{
			name:  "magnitude of y unit vector",
			tuple: MakeVector(0, 1, 0),
			want:  1,
		},
		{
			name:  "magnitude of z unit vector",
			tuple: MakeVector(0, 0, 1),
			want:  1,
		},
		{
			name:  "magnitude of arbitrary vector",
			tuple: MakeVector(1, 2, 3),
			want:  math.Sqrt(14),
		},
		{
			name:  "magnitude of arbitrary negative vector",
			tuple: MakeVector(-1, -2, -3),
			want:  math.Sqrt(14),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			if got := tt.tuple.Magnitude(); got != tt.want {
				t1.Errorf("Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Multiply(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		scale float64
		want  Tuple
	}{
		{
			name:  "multiply by scalar",
			tuple: Tuple{1, -2, 3, -4},
			scale: 3.5,
			want:  Tuple{3.5, -7, 10.5, -14},
		},
		{
			name:  "multiply by fraction",
			tuple: Tuple{1, -2, 3, -4},
			scale: 0.5,
			want:  Tuple{0.5, -1, 1.5, -2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			if got := tt.tuple.Multiply(tt.scale); got != tt.want {
				t1.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Negate(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		want  Tuple
	}{
		{
			name:  "negate arbitrary tuple",
			tuple: Tuple{1, -2, 3, -4},
			want:  Tuple{-1, 2, -3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			if got := tt.tuple.Negate(); got != tt.want {
				t1.Errorf("Negate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Normalized(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		want  Tuple
	}{
		{
			name:  "single axis vector",
			tuple: MakeVector(4, 0, 0),
			want:  MakeVector(1, 0, 0),
		},
		{
			name:  "arbitrary vector",
			tuple: MakeVector(1, 2, 3),
			want: MakeVector(
				1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			got := tt.tuple.Normalized()
			if got != tt.want {
				t1.Errorf("Normalized() = %v, want %v", got, tt.want)
			}

			if !Float64Equal(got.Magnitude(), 1) {
				t1.Errorf(
					"Length of normalized vector = %v, want 1",
					got.Magnitude(),
				)
			}
		})
	}
}

func TestTuple_Subtract(t *testing.T) {
	tests := []tupleOperationTest{
		{
			"point subtraction",
			MakePoint(3, 2, 1),
			MakePoint(5, 6, 7),
			MakeVector(-2, -4, -6),
		},
		{
			"subtract vector from point",
			MakePoint(3, 2, 1),
			MakeVector(5, 6, 7),
			MakePoint(-2, -4, -6),
		},
		{
			"vector subtraction",
			MakeVector(3, 2, 1),
			MakeVector(5, 6, 7),
			MakeVector(-2, -4, -6),
		},
		{
			"vector subtraction from 0",
			MakeVector(0, 0, 0),
			MakeVector(1, -2, 3),
			MakeVector(-1, 2, -3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			if got := tt.base.Subtract(tt.other); got != tt.want {
				t1.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakePoint(t *testing.T) {
	x := 4.3
	y := -4.2
	z := 3.1

	point := MakePoint(x, y, z)

	if point.X != x {
		t.Errorf("point.X = %v, want %v", point.X, x)
	}

	if point.Y != y {
		t.Errorf("point.Y = %v, want %v", point.Y, y)
	}

	if point.Z != z {
		t.Errorf("point.Z = %v, want %v", point.Z, z)
	}

	if point.W != 1 {
		t.Errorf("point.W = %v, want 1", point.W)
	}
}

func TestMakeVector(t *testing.T) {
	x := 4.3
	y := -4.2
	z := 3.1

	vector := MakeVector(x, y, z)

	if vector.X != x {
		t.Errorf("vector.X = %v, want %v", vector.X, x)
	}

	if vector.Y != y {
		t.Errorf("vector.Y = %v, want %v", vector.Y, y)
	}

	if vector.Z != z {
		t.Errorf("vector.Z = %v, want %v", vector.Z, z)
	}

	if vector.W != 0 {
		t.Errorf("vector.W = %v, want 0", vector.W)
	}
}
