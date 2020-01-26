package main

import "testing"

func TestMakeMaterial(t *testing.T) {
	m := MakeMaterial()

	if !m.Color.Equals(MakeColor(1, 1, 1)) {
		t.Errorf("Expected default color to be %v, got %v", MakeColor(1, 1, 1), m.Color)
	}

	if m.Ambient != 0.1 {
		t.Errorf("Expected default ambient value to be %v, got %v", 0.1, m.Ambient)
	}

	if m.Diffuse != 0.9 {
		t.Errorf("Expected default diffuse value to be %v, got %v", 0.9, m.Diffuse)
	}

	if m.Specular != 0.9 {
		t.Errorf("Expected default specular value to be %v, got %v", 0.9, m.Specular)
	}

	if m.Shininess != 200.0 {
		t.Errorf("Expected default shininess value to be %v, got %v", 200.0, m.Shininess)
	}
}
