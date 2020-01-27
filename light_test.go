package main

import (
	"math"
	"testing"
)

func TestMakePointLight(t *testing.T) {
	intensity := MakeColor(1, 1, 1)
	position := MakePoint(0, 0, 0)

	light := MakePointLight(position, intensity)

	if got := light.Position; !got.Equals(position) {
		t.Errorf("Expected light position to be %v, got %v", position, got)
	}

	if got := light.Intensity; !got.Equals(intensity) {
		t.Errorf("Expected light intensity to be %v, got %v", intensity, got)
	}
}

func TestLighting(t *testing.T) {
	testCases := []struct {
		name      string
		eyeVector Tuple
		normal    Tuple
		light     PointLight
		want      Color
	}{
		{
			"eye between light and surface",
			MakeVector(0, 0, -1),
			MakeVector(0, 0, -1),
			MakePointLight(MakePoint(0, 0, -10), MakeColor(1, 1, 1)),
			MakeColor(1.9, 1.9, 1.9),
		},
		{
			"eye between light and surface - eye offset 45 degrees",
			MakeVector(0, math.Sqrt2/2, -math.Sqrt2/2),
			MakeVector(0, 0, -1),
			MakePointLight(MakePoint(0, 0, -10), MakeColor(1, 1, 1)),
			MakeColor(1, 1, 1),
		},
		{
			"eye opposite surface - light offset 45 degrees",
			MakeVector(0, 0, -1),
			MakeVector(0, 0, -1),
			MakePointLight(MakePoint(0, 10, -10), MakeColor(1, 1, 1)),
			MakeColor(0.7364, 0.7364, 0.7364),
		},
		{
			"eye in path of reflection vector",
			MakeVector(0, -math.Sqrt2/2, -math.Sqrt2/2),
			MakeVector(0, 0, -1),
			MakePointLight(MakePoint(0, 10, -10), MakeColor(1, 1, 1)),
			MakeColor(1.6364, 1.6364, 1.6364),
		},
		{
			"light behind surface",
			MakeVector(0, 0, -1),
			MakeVector(0, 0, -1),
			MakePointLight(MakePoint(0, 0, 10), MakeColor(1, 1, 1)),
			MakeColor(0.1, 0.1, 0.1),
		},
	}
	for _, tt := range testCases {
		m := MakeMaterial()
		position := MakePoint(0, 0, 0)
		t.Run(tt.name, func(t *testing.T) {
			if got := Lighting(m, tt.light, position, tt.eyeVector, tt.normal); !got.Equals(tt.want) {
				t.Errorf("Expected lighting to produce color %v, got %v", tt.want, got)
			}
		})
	}
}
