package main

import "math"

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

// Get the color of a position given a material, light source, observer, and the
// normal of the illuminated surface.
func Lighting(material Material, light PointLight, position Tuple, eyeVector Tuple, normal Tuple) Color {
	// Initial color is a combination of the material's color and the light's
	// color.
	effectiveColor := material.Color.Blend(light.Intensity)

	lightVector := light.Position.Subtract(position).Normalized()

	// The ambient color is the color contribution from "background" light or
	// the color shown with no light sources.
	ambient := effectiveColor.Multiply(material.Ambient)

	diffuse := MakeColor(0, 0, 0)
	specular := MakeColor(0, 0, 0)
	// The dot product of the vector to the light source and the normal vector
	// represents the cosine of the angle between the two. We only need to
	// compute the diffuse and specular colors if the cosine is positive,
	// because a negative result indicates the light source is on the other side
	// of the surface.
	lightDotNormal := lightVector.Dot(normal)
	if lightDotNormal >= 0 {
		diffuse = effectiveColor.Multiply(material.Diffuse).Multiply(lightDotNormal)

		// The following represents the cosine of the angle between the
		// reflection vector and they eye vector. A negative number means the
		// light reflects away from the eye.
		reflectionVector := Reflect(lightVector.Multiply(-1), normal)
		reflectionDotEye := reflectionVector.Dot(eyeVector)
		if reflectionDotEye > 0 {
			factor := math.Pow(reflectionDotEye, material.Shininess)
			specular = light.Intensity.Multiply(material.Specular).Multiply(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
