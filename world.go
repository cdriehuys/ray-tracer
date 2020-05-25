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
					Color:    MakeColor(0.8, 1, 0.6),
					Diffuse:  0.7,
					Specular: 0.2,
				},
				Transform: IdentityMatrix4,
			},
			MakeSphereTransformed(MakeScale(0.5, 0.5, 0.5)),
		},
	}
}

func (w World) Intersect(ray Ray) (intersections Intersections) {
	for _, object := range w.Objects {
		intersections = append(intersections, object.Intersect(ray)...)
	}

	if intersections != nil {
		intersections.Sort()
	}

	return intersections
}
