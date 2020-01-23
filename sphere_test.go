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
		want Intersections
	}{
		{
			"intersects at two points",
			MakeRay(MakePoint(0, 0, -5), MakeVector(0, 0, 1)),
			Intersections{
				MakeIntersection(4, sphere),
				MakeIntersection(6, sphere),
			},
		},
		{
			"intersects at tangent",
			MakeRay(MakePoint(0, 1, -5), MakeVector(0, 0, 1)),
			Intersections{
				MakeIntersection(5, sphere),
				MakeIntersection(5, sphere),
			},
		},
		{
			"miss",
			MakeRay(MakePoint(0, 2, -5), MakeVector(0, 0, 1)),
			Intersections{},
		},
		{
			"ray inside sphere",
			MakeRay(MakePoint(0, 0, 0), MakeVector(0, 0, 1)),
			Intersections{
				MakeIntersection(-1, sphere),
				MakeIntersection(1, sphere),
			},
		},
		{
			"sphere behind ray",
			MakeRay(MakePoint(0, 0, 5), MakeVector(0, 0, 1)),
			Intersections{
				MakeIntersection(-6, sphere),
				MakeIntersection(-4, sphere),
			},
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
