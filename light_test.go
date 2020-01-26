package main

import "testing"

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
