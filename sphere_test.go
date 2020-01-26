package main

import (
	"reflect"
	"testing"
)

func TestMakeSphere(t *testing.T) {
	sphere := MakeSphere()

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
