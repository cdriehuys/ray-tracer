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
