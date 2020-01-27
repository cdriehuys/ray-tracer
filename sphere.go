package main

import "math"

type Sphere struct {
	material  Material
	Transform Matrix
}

func MakeSphere() Sphere {
	return Sphere{
		material:  MakeMaterial(),
		Transform: IdentityMatrix4,
	}
}

func MakeSphereTransformed(transform Matrix) Sphere {
	return Sphere{
		material:  MakeMaterial(),
		Transform: transform,
	}
}

// Get the values of t at which the given ray intersects the sphere.
func (s Sphere) Intersect(ray Ray) Intersections {
	// Apply the sphere's transformations by applying their inverse to the ray.
	ray = ray.Transform(s.Transform.Inverted())

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

// Get the material used by the sphere.
func (s Sphere) Material() Material {
	return s.material
}

// Get the normal vector at a point on the surface of a sphere. This point is
// given in world space (as opposed to object space).
func (s Sphere) NormalAt(worldPoint Tuple) Tuple {
	objectPoint := s.Transform.Inverted().TupleMultiply(worldPoint)
	// We're subtracting the origin of the sphere, which is always the origin in
	// object space.
	objectNormal := objectPoint.Subtract(MakePoint(0, 0, 0))
	worldNormal := s.Transform.Inverted().Transposed().TupleMultiply(objectNormal)
	// Since we should have ignored the 4th row and column of the matrix in the
	// computation above, the 4th row (which includes w for our tuple) may have
	// been messed with. To compensate for this, we manually set w to 0, which
	// represents a vector.
	worldNormal.W = 0

	return worldNormal.Normalized()
}
