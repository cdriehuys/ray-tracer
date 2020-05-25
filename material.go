package main

type Material struct {
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func MakeMaterial() Material {
	return Material{
		Color:     MakeColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}

// Determine if one material is equivalent to another.
func (mat Material) Equals(other Material) bool {
	return mat.Color.Equals(other.Color) &&
		Float64Equal(mat.Ambient, other.Ambient) &&
		Float64Equal(mat.Diffuse, other.Diffuse) &&
		Float64Equal(mat.Specular, other.Specular) &&
		Float64Equal(mat.Shininess, other.Shininess)
}
