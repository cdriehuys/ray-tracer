package main

import (
	"reflect"
	"testing"
)

func TestSphere_Intersect(t *testing.T) {
	sphere := Sphere{}
	testCases := []struct {
		name string
		ray  Ray
		want []float64
	}{
		{
			"intersects at two points",
			MakeRay(MakePoint(0, 0, -5), MakeVector(0, 0, 1)),
			[]float64{4.0, 6.0},
		},
		{
			"intersects at tangent",
			MakeRay(MakePoint(0, 1, -5), MakeVector(0, 0, 1)),
			[]float64{5.0, 5.0},
		},
		{
			"miss",
			MakeRay(MakePoint(0, 2, -5), MakeVector(0, 0, 1)),
			[]float64{},
		},
		{
			"ray inside sphere",
			MakeRay(MakePoint(0, 0, 0), MakeVector(0, 0, 1)),
			[]float64{-1.0, 1.0},
		},
		{
			"sphere behind ray",
			MakeRay(MakePoint(0, 0, 5), MakeVector(0, 0, 1)),
			[]float64{-6.0, -4.0},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := sphere.Intersect(tt.ray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect did not produce expected results:\nExpected: %v\nReceived: %v", tt.want, got)
			}
		})
	}
}
