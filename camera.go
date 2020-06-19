package main

import "math"

// A camera represents a view into a world.
type Camera struct {
	// The width of the camera's view in pixels.
	Width int
	// The height of the camera's view in pixels.
	Height int

	// The angular width of what the camera sees specified in radians.
	FieldOfView float64

	// A matrix indicating how the *world* is oriented with respect to the
	// camera.
	Transform Matrix

	// Half the width of the view in world-space units.
	halfWidth float64
	// Half the height of the view in world-space units.
	halfHeight float64
	// The size of a pixel in world-space units.
	pixelSize float64
}

func MakeCamera(width, height int, fov float64) Camera {
	camera := Camera{
		Width:       width,
		Height:      height,
		FieldOfView: fov,
		Transform:   IdentityMatrix4,
	}
	camera.computeCameraPixelSize()

	return camera
}

func (c Camera) MakeRayForPixel(x, y int) Ray {
	// Compute offsets from the edge of the canvas to the pixel's center.
	offsetX := (float64(x) + 0.5) * c.pixelSize
	offsetY := (float64(y) + 0.5) * c.pixelSize

	// The untransformed coordinates of the pixel in world-space.
	worldX := c.halfWidth - offsetX
	worldY := c.halfHeight - offsetY

	// Transform the canvas point and origin, remembering that the canvas is at
	// z = -1.
	inverseCameraTransform := c.Transform.Inverted()
	pixel := inverseCameraTransform.TupleMultiply(MakePoint(worldX, worldY, -1))
	origin := inverseCameraTransform.TupleMultiply(MakePoint(0, 0, 0))
	direction := pixel.Subtract(origin).Normalized()

	return MakeRay(origin, direction)
}

// Compute the size of a pixel in world-space units for a camera with the given
// width and height. The camera assumes that the canvas is exactly one world-
// space unit away from the camera.
func (c *Camera) computeCameraPixelSize() {
	// Compute the size of half the view in world-space units. This can
	// represent either half the view's height or half the view's width
	// depending on the aspect ratio of the camera. Think of
	halfView := math.Tan(c.FieldOfView / 2)
	aspectRatio := float64(c.Width) / float64(c.Height)

	if aspectRatio >= 1 {
		// Horizontal canvas
		c.halfWidth = halfView
		c.halfHeight = halfView / aspectRatio
	} else {
		// Vertical canvas
		c.halfWidth = halfView * aspectRatio
		c.halfHeight = halfView
	}

	c.pixelSize = c.halfWidth * 2 / float64(c.Width)
}
