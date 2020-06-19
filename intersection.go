package main

import "sort"

type Intersection struct {
	T      float64
	Object Object
}

func MakeIntersection(t float64, object Object) Intersection {
	return Intersection{t, object}
}

// Prepare some useful properties about the intersection for later use.
func (i Intersection) PrepareComputations(ray Ray) IntersectionComputation {
	intersectionPoint := ray.Position(i.T)
	eyeVector := ray.Direction.Negate()
	normalVector := i.Object.NormalAt(intersectionPoint)

	// The hit occurred on the inside of the shape if the eye vector and normal
	// vector are pointing roughly in opposite directions.
	inside := normalVector.Dot(eyeVector) < 0
	if inside {
		normalVector = normalVector.Negate()
	}

	return IntersectionComputation{
		T:            i.T,
		Object:       i.Object,
		Inside:       inside,
		Point:        intersectionPoint,
		EyeVector:    eyeVector,
		NormalVector: normalVector,
	}
}

type Intersections []Intersection

// Get the intersection in the collection with the lowest, non-negative t value.
// If no such value exists, the boolean return value will be false and the
// returned intersection should be ignored.
func (i Intersections) Hit() (Intersection, bool) {
	i.Sort()

	for _, intersection := range i {
		if intersection.T >= 0 {
			return intersection, true
		}
	}

	return Intersection{}, false
}

// Sort the intersections in place by t-value, ascending.
func (i Intersections) Sort() {
	sort.Slice(i, func(x, y int) bool {
		return i[x].T < i[y].T
	})
}
