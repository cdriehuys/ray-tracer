package main

import (
	"math"
	"reflect"
	"testing"
)

var sqrt3 = math.Sqrt(3)

func TestMakeSphere(t *testing.T) {
	sphere := MakeSphere()

	if got := sphere.Material; !reflect.DeepEqual(got, MakeMaterial()) {
		t.Errorf("Expected sphere to have default material, got %v", got)
	}

	if got := sphere.Transform; !got.Equals(IdentityMatrix4) {
		t.Errorf("Expected default transform to be the identity matrix, got %v", got)
	}
}

func TestMakeSphereTransformed(t *testing.T) {
	transform := MakeTranslation(1, 2, 3)

	sphere := MakeSphereTransformed(transform)

	if got := sphere.Transform; !got.Equals(transform) {
		t.Errorf("Expected sphere's transform to be %v, got %v", transform, got)
	}
}

func TestSphere_Intersect(t *testing.T) {
	testCases := []struct {
		name   string
		ray    Ray
		sphere Sphere
		want   []float64
	}{
		{
			"intersects at two points",
			MakeRay(MakePoint(0, 0, -5), MakeVector(0, 0, 1)),
			MakeSphere(),
			[]float64{4, 6},
		},
		{
			"intersects at tangent",
			MakeRay(MakePoint(0, 1, -5), MakeVector(0, 0, 1)),
			MakeSphere(),
			[]float64{5, 5},
		},
		{
			"miss",
			MakeRay(MakePoint(0, 2, -5), MakeVector(0, 0, 1)),
			MakeSphere(),
			[]float64{},
		},
		{
			"ray inside sphere",
			MakeRay(MakePoint(0, 0, 0), MakeVector(0, 0, 1)),
			MakeSphere(),
			[]float64{-1, 1},
		},
		{
			"sphere behind ray",
			MakeRay(MakePoint(0, 0, 5), MakeVector(0, 0, 1)),
			MakeSphere(),
			[]float64{-6, -4},
		},
		{
			"scaled sphere",
			MakeRay(MakePoint(0, 0, -5), MakeVector(0, 0, 1)),
			MakeSphereTransformed(MakeScale(2, 2, 2)),
			[]float64{3, 7},
		},
		{
			"translated sphere",
			MakeRay(MakePoint(0, 0, -5), MakeVector(0, 0, 1)),
			MakeSphereTransformed(MakeTranslation(5, 0, 0)),
			[]float64{},
		},
	}
	for _, tt := range testCases {
		wantIntersections := Intersections{}
		for _, t := range tt.want {
			wantIntersections = append(wantIntersections, MakeIntersection(t, tt.sphere))
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sphere.Intersect(tt.ray); !reflect.DeepEqual(got, wantIntersections) {
				t.Errorf("Intersect did not produce expected results:\nExpected: %v\nReceived: %v", wantIntersections, got)
			}
		})
	}
}

func TestSphere_NormalAt(t *testing.T) {
	testCases := []struct {
		name   string
		sphere Sphere
		point  Tuple
		want   Tuple
	}{
		{
			"x-axis",
			MakeSphere(),
			MakePoint(1, 0, 0),
			MakeVector(1, 0, 0),
		},
		{
			"y-axis",
			MakeSphere(),
			MakePoint(0, 1, 0),
			MakeVector(0, 1, 0),
		},
		{
			"z-axis",
			MakeSphere(),
			MakePoint(0, 0, 1),
			MakeVector(0, 0, 1),
		},
		{
			"nonaxial point",
			MakeSphere(),
			MakePoint(sqrt3/3, sqrt3/3, sqrt3/3),
			MakeVector(sqrt3/3, sqrt3/3, sqrt3/3),
		},
		{
			"translated sphere",
			MakeSphereTransformed(MakeTranslation(0, 1, 0)),
			MakePoint(0, 1.70711, -0.70711),
			MakeVector(0, 0.70711, -0.70711),
		},
		{
			"transformed sphere",
			MakeSphereTransformed(
				MakeScale(1, 0.5, 1).Multiply(MakeZRotation(math.Pi / 5)),
			),
			MakePoint(0, math.Sqrt2/2, -math.Sqrt2/2),
			MakeVector(0, 0.97014, -0.24254),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			normal := tt.sphere.NormalAt(tt.point)
			if !normal.Equals(tt.want) {
				t.Errorf("Expected normal to be %v, got %v", tt.want, normal)
			}

			if !normal.Equals(normal.Normalized()) {
				t.Errorf(
					"Expected normal to be normalized but it was %v (magitude %v)",
					normal,
					normal.Magnitude(),
				)
			}
		})
	}
}
