package main

import "testing"

func TestTuple_IsPoint(t1 *testing.T) {
	tests := []struct {
		name   string
		tuple Tuple
		want   bool
	}{
		{
			name: "point",
			tuple: MakePoint(4.3, -4.2, 3.1),
			want: true,
		},
		{
			name: "vector",
			tuple: MakeVector(4.3, -4.2, 3.1),
			want: false,
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
		name   string
		tuple Tuple
		want   bool
	}{
		{
			name: "point",
			tuple: MakePoint(4.3, -4.2, 3.1),
			want: false,
		},
		{
			name: "vector",
			tuple: MakeVector(4.3, -4.2, 3.1),
			want: true,
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
