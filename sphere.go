package main

import "math"

type Sphere struct{}

// Get the values of t at which the given ray intersects the sphere.
func (s Sphere) Intersect(ray Ray) Intersections {
	// Assume the sphere is at the origin.
	sphereToRay := ray.Origin.Subtract(MakePoint(0, 0, 0))

	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - (4 * a * c)

	if discriminant < 0 {
		return Intersections{}
	}

	discriminantRoot := math.Sqrt(discriminant)
	t1 := (-b - discriminantRoot) / (2 * a)
	t2 := (-b + discriminantRoot) / (2 * a)

	return Intersections{
		MakeIntersection(t1, s),
		MakeIntersection(t2, s),
	}
}