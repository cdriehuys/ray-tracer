package main

import "math"

type Sphere struct{}

func (s Sphere) Intersect(ray Ray) []float64 {
	sphereToRay := ray.Origin.Subtract(MakePoint(0, 0, 0))

	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - (4 * a * c)

	if discriminant < 0 {
		return []float64{}
	}

	discriminantRoot := math.Sqrt(discriminant)
	t1 := (-b - discriminantRoot) / (2 * a)
	t2 := (-b + discriminantRoot) / (2 * a)

	return []float64{t1, t2}
}
