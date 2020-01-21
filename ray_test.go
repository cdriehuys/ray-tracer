package main

import (
	"fmt"
	"testing"
)

func TestMakeRay(t *testing.T) {
	origin := MakePoint(1, 2, 3)
	direction := MakeVector(4, 5, 6)

	ray := MakeRay(origin, direction)

	if got := ray.Origin; !got.Equals(origin) {
		t.Errorf("Expected ray.Origin = %v, got %v", origin, got)
	}

	if got := ray.Direction; !got.Equals(direction) {
		t.Errorf("Expected ray.Direction = %v, got %v", direction, got)
	}
}

func TestRay_Position(t *testing.T) {
	origin := MakePoint(2, 3, 4)
	direction := MakeVector(1, 0, 0)
	ray := MakeRay(origin, direction)

	testCases := []struct {
		time float64
		want Tuple
	}{
		{0, MakePoint(2, 3, 4)},
		{1, MakePoint(3, 3, 4)},
		{-1, MakePoint(1, 3, 4)},
		{2.5, MakePoint(4.5, 3, 4)},
	}
	for _, tt := range testCases {
		t.Run(fmt.Sprintf("t=%v", tt.time), func(t *testing.T) {
			if got := ray.Position(tt.time); !got.Equals(tt.want) {
				t.Errorf("Expected position(%v) = %v, got %v", tt.time, tt.want, got)
			}
		})
	}
}
