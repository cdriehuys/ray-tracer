package main

import "sort"

type Shape interface {
	Intersect(Ray) Intersections
}

type Intersection struct {
	T      float64
	Object Shape
}

func MakeIntersection(t float64, object Shape) Intersection {
	return Intersection{t, object}
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
