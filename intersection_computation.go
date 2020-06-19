package main

// Data structure containing precomputed properties for a specific intersection.
type IntersectionComputation struct {
	// The t-value of the intersection that produced this computation.
	T float64
	// The object of the intersection that produced this computation.
	Object Object

	// Boolean indicating if the hit occurred on the inside of the shape.
	Inside bool
	// The point where the intersection occurred.
	Point Tuple
	// A vector pointing from the intersection point back to the observer's eye.
	EyeVector Tuple
	// The normal vector of the intersected object at the point of intersection.
	NormalVector Tuple
}
