package main

type PointLight struct {
	Position  Tuple
	Intensity Color
}

func MakePointLight(position Tuple, intensity Color) PointLight {
	return PointLight{
		Position:  position,
		Intensity: intensity,
	}
}
