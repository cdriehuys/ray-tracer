package main

// A world stores the objects and light sources that make up a scene.
type World struct {
	// The light source used to illuminate the world.
	Light PointLight

	// The renderable objects in the world.
	Objects []Object
}

// Create an empty world.
func MakeWorld() World {
	return World{}
}

// Create a world with a default light source and objects.
func MakeDefaultWorld() World {
	return World{
		Light: MakePointLight(
			MakePoint(-10, 10, -10),
			MakeColor(1, 1, 1),
		),
		Objects: []Object{
			Sphere{
				material: Material{
					Color:     MakeColor(0.8, 1, 0.6),
					Ambient:   0.1,
					Diffuse:   0.7,
					Specular:  0.2,
					Shininess: 200.0,
				},
				transform: IdentityMatrix4,
			},
			MakeSphereTransformed(MakeScale(0.5, 0.5, 0.5)),
		},
	}
}

// Compute the color resulting from the given ray intersecting the objects in
// the world.
func (w World) ColorAt(ray Ray) Color {
	intersections := w.intersect(ray)
	intersection, hit := intersections.Hit()

	// No hit means we should return the color of the void.
	if !hit {
		return MakeColor(0, 0, 0)
	}

	intersectionComps := intersection.PrepareComputations(ray)

	return w.shadeHit(intersectionComps)
}

func (w World) intersect(ray Ray) (intersections Intersections) {
	for _, object := range w.Objects {
		intersections = append(intersections, object.Intersect(ray)...)
	}

	if intersections != nil {
		intersections.Sort()
	}

	return intersections
}

// Find the color that should be produced at the location of the given
// intersection.
func (w World) shadeHit(computation IntersectionComputation) Color {
	return Lighting(
		computation.Object.Material(),
		w.Light,
		computation.Point,
		computation.EyeVector,
		computation.NormalVector,
	)
}
