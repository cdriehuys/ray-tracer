package main

import (
	"math"
	"testing"
)

func TestRender(t *testing.T) {
	world := MakeDefaultWorld()
	camera := MakeCamera(11, 11, math.Pi/2)
	from := MakePoint(0, 0, -5)
	to := MakePoint(0, 0, 0)
	up := MakeVector(0, 1, 0)
	camera.Transform = ViewTransform(from, to, up)

	want := MakeColor(0.38066, 0.47583, 0.2855)

	image := Render(camera, world)

	if got := image.GetPixel(5, 5); !want.Equals(got) {
		t.Errorf("Expected pixel at (5, 5) to be %v, got %v", want, got)
	}
}
