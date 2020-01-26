package main

import (
	"fmt"
	"math"
	"reflect"
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

func TestRay_Transform(t *testing.T) {
	testCases := []struct {
		name      string
		ray       Ray
		transform Matrix
		want      Ray
	}{
		{
			"translation",
			MakeRay(MakePoint(1, 2, 3), MakeVector(0, 1, 0)),
			MakeTranslation(3, 4, 5),
			MakeRay(MakePoint(4, 6, 8), MakeVector(0, 1, 0)),
		},
		{
			"scale",
			MakeRay(MakePoint(1, 2, 3), MakeVector(0, 1, 0)),
			MakeScale(2, 3, 4),
			MakeRay(MakePoint(2, 6, 12), MakeVector(0, 3, 0)),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ray.Transform(tt.transform); !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"Expected different ray transformation:\nRay: %v\nTransform: %v\nExpected: %v\nReceived: %v",
					tt.ray,
					tt.transform,
					tt.want,
					got,
				)
			}
		})
	}
}

func TestReflect(t *testing.T) {
	testCases := []struct {
		name   string
		in     Tuple
		normal Tuple
		want   Tuple
	}{
		{
			"45 degrees",
			MakeVector(1, -1, 0),
			MakeVector(0, 1, 0),
			MakeVector(1, 1, 0),
		},
		{
			"slanted surface",
			MakeVector(0, -1, 0),
			MakeVector(math.Sqrt2/2, math.Sqrt2/2, 0),
			MakeVector(1, 0, 0),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reflect(tt.in, tt.normal); !got.Equals(tt.want) {
				t.Errorf("Expected reflecting %v around %v to be %v, got %v", tt.in, tt.normal, tt.want, got)
			}
		})
	}
}
