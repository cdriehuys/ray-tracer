package main

// An object is any structure that has the attributes required to be rendered.
type Object interface {
	// Find the intersections that the object has with a specific ray.
	Intersect(Ray) Intersections

	// Get the material that determines how the object reflects light.
	Material() Material

	// Get the normal vector at any point on the object's surface. The point on
	// the object's surface is given in world space (as opposed to object
	// space).
	NormalAt(Tuple) Tuple
}
