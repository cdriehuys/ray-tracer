package main

import "math"

// Create a matrix representing a scale transformation. The X, Y, and Z axis may
// all be scaled by different factors.
func MakeScale(x, y, z float64) Matrix {
	return MakeMatrix4(
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	)
}

// Create a matrix representing a shear transformation where components by
// scaled in proportion to the other axis.
func MakeShear(xToY, xToZ, yToX, yToZ, zToX, zToY float64) Matrix {
	return MakeMatrix4(
		1, xToY, xToZ, 0,
		yToX, 1, yToZ, 0,
		zToX, zToY, 1, 0,
		0, 0, 0, 1,
	)
}

// Create a matrix representing a translation transformation by the given
// amounts on the X, Y, and Z axis.
func MakeTranslation(x, y, z float64) Matrix {
	return MakeMatrix4(
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	)
}

// Create a matrix representing a rotation around the x-axis by the specified
// number of radians.
func MakeXRotation(radians float64) Matrix {
	return MakeMatrix4(
		1, 0, 0, 0,
		0, math.Cos(radians), -math.Sin(radians), 0,
		0, math.Sin(radians), math.Cos(radians), 0,
		0, 0, 0, 1,
	)
}

// Create a matrix representing a rotation around the y-axis by the specified
// number of radians.
func MakeYRotation(radians float64) Matrix {
	return MakeMatrix4(
		math.Cos(radians), 0, math.Sin(radians), 0,
		0, 1, 0, 0,
		-math.Sin(radians), 0, math.Cos(radians), 0,
		0, 0, 0, 1,
	)
}

// Create a matrix representing a rotation around the z-axis by the specified
// number of radians.
func MakeZRotation(radians float64) Matrix {
	return MakeMatrix4(
		math.Cos(radians), -math.Sin(radians), 0, 0,
		math.Sin(radians), math.Cos(radians), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}
