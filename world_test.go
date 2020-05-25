package main

import (
	"reflect"
	"testing"
)

func TestMakeWorld(t *testing.T) {
	expected := World{}

	if got := MakeWorld(); !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected an empty world; got %v", got)
	}
}

func TestMakeDefaultWorld(t *testing.T) {
	expectedLight := MakePointLight(
		MakePoint(-10, 10, -10),
		MakeColor(1, 1, 1),
	)
	expectedObjects := []Object{
		Sphere{
			material: Material{
				Color:    MakeColor(0.8, 1, 0.6),
				Diffuse:  0.7,
				Specular: 0.2,
			},
			Transform: IdentityMatrix4,
		},
		MakeSphereTransformed(MakeScale(0.5, 0.5, 0.5)),
	}

	world := MakeDefaultWorld()
	if got := world.Light; !reflect.DeepEqual(expectedLight, got) {
		t.Errorf("World light does not match the expected light:\nExpected: %v\nReceived: %v", expectedLight, got)
	}

	for _, expectedObj := range expectedObjects {
		assertContainsObject(t, world.Objects, expectedObj)
	}
}

func TestWorld_Intersect(t *testing.T) {
	world := MakeDefaultWorld()
	ray := MakeRay(MakePoint(0, 0, -5), MakeVector(0, 0, 1))
	wantIntersections := []float64{4, 4.5, 5.5, 6}

	intersections := world.Intersect(ray)
	if got := len(intersections); got != len(wantIntersections) {
		t.Errorf("Expected %d intersection(s); got %d", len(wantIntersections), got)
	}

	for i, intersection := range intersections {
		if got := intersection.T; !Float64Equal(wantIntersections[i], got) {
			t.Errorf("Expected intersection %d to have t-value %f; got %f", i, wantIntersections[i], got)
		}
	}
}

func assertContainsObject(t *testing.T, objects []Object, want Object) {
	for _, obj := range objects {
		if objectsEqual(obj, want) {
			return
		}
	}

	t.Errorf("Could not find object %v", want)
}

func objectsEqual(a, b Object) bool {
	return a.Material().Equals(b.Material())
}
