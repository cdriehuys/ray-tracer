package main

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

func MakeRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

// Get the position of a ray at the given time.
func (r Ray) Position(t float64) Tuple {
	return r.Origin.Add(r.Direction.Multiply(t))
}

// Create a new ray by applying a transformation to the current ray.
func (r Ray) Transform(transform Matrix) Ray {
	return MakeRay(
		transform.TupleMultiply(r.Origin),
		transform.TupleMultiply(r.Direction),
	)
}

// Reflect a ray around a normal vector.
func Reflect(in, normal Tuple) Tuple {
	return in.Subtract(normal.Multiply(2 * in.Dot(normal)))
}
