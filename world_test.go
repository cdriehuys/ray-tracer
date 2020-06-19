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
				Color:     MakeColor(0.8, 1, 0.6),
				Ambient:   0.1,
				Diffuse:   0.7,
				Specular:  0.2,
				Shininess: 200,
			},
			transform: IdentityMatrix4,
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

func TestWorld_ColorAt(t *testing.T) {
	testCases := []struct {
		name  string
		world World
		ray   Ray
		want  Color
	}{
		{
			"ray misses",
			MakeDefaultWorld(),
			MakeRay(MakePoint(0, 0, -5), MakeVector(0, 1, 0)),
			MakeColor(0, 0, 0),
		},
		{
			"ray hit in default world",
			MakeDefaultWorld(),
			MakeRay(MakePoint(0, 0, -5), MakeVector(0, 0, 1)),
			MakeColor(0.38066, 0.47583, 0.2855),
		},
		{
			"intersection behind ray",
			func() World {
				world := MakeDefaultWorld()
				outer := world.Objects[0].(Sphere)
				outer.material.Ambient = 1
				world.Objects[0] = outer

				inner := world.Objects[1].(Sphere)
				inner.material.Ambient = 1
				world.Objects[1] = inner

				return world
			}(),
			MakeRay(MakePoint(0, 0, 0.75), MakeVector(0, 0, -1)),
			MakeDefaultWorld().Objects[1].Material().Color,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.world.ColorAt(tt.ray); !tt.want.Equals(got) {
				t.Errorf("Expected color %v, got %v", tt.want, got)
			}
		})
	}
}

func TestWorld_Intersect(t *testing.T) {
	world := MakeDefaultWorld()
	ray := MakeRay(MakePoint(0, 0, -5), MakeVector(0, 0, 1))
	wantIntersections := []float64{4, 4.5, 5.5, 6}

	intersections := world.intersect(ray)
	if got := len(intersections); got != len(wantIntersections) {
		t.Errorf("Expected %d intersection(s); got %d", len(wantIntersections), got)
	}

	for i, intersection := range intersections {
		if got := intersection.T; !Float64Equal(wantIntersections[i], got) {
			t.Errorf("Expected intersection %d to have t-value %f; got %f", i, wantIntersections[i], got)
		}
	}
}

func TestWorld_ShadeHit(t *testing.T) {
	defaultWorld := MakeDefaultWorld()

	testCases := []struct {
		name         string
		light        PointLight
		ray          Ray
		intersection Intersection
		want         Color
	}{
		{
			"outside intersection",
			defaultWorld.Light,
			MakeRay(MakePoint(0, 0, -5), MakeVector(0, 0, 1)),
			MakeIntersection(4, defaultWorld.Objects[0]),
			MakeColor(0.38066, 0.47583, 0.2855),
		},
		{
			"inside intersection",
			MakePointLight(MakePoint(0, 0.25, 0), MakeColor(1, 1, 1)),
			MakeRay(MakePoint(0, 0, 0), MakeVector(0, 0, 1)),
			MakeIntersection(0.5, defaultWorld.Objects[1]),
			MakeColor(0.90498, 0.90498, 0.90498),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			world := MakeDefaultWorld()
			world.Light = tt.light
			comps := tt.intersection.PrepareComputations(tt.ray)

			if got := world.shadeHit(comps); !tt.want.Equals(got) {
				t.Errorf("Expected color of hit to be %v; got %v", tt.want, got)
			}
		})
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
	return a.Material().Equals(b.Material()) &&
		a.Transform().Equals(b.Transform())
}
