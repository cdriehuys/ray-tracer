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

func TestMaterial_Equals(t *testing.T) {
	testCases := []struct {
		name      string
		materialA Material
		materialB Material
		want      bool
	}{
		{
			"different colors",
			Material{Color: MakeColor(1, 1, 0)},
			Material{Color: MakeColor(0, 0, 1)},
			false,
		},
		{
			"different ambient values",
			Material{Ambient: 12},
			Material{Ambient: 42},
			false,
		},
		{
			"different diffuse values",
			Material{Diffuse: 13},
			Material{Diffuse: 76},
			false,
		},
		{
			"different specular values",
			Material{Specular: 7},
			Material{Specular: 13},
			false,
		},
		{
			"different shininess values",
			Material{Shininess: 8675309},
			Material{Shininess: 9001},
			false,
		},
		{
			"same material",
			MakeMaterial(),
			MakeMaterial(),
			true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.materialA.Equals(tt.materialB); got != tt.want {
				t.Errorf("Expected equality to evaluate to %v\nExpected: %v\nReceived: %v\n", tt.want, tt.materialA, tt.materialB)
			}
		})
	}
}
