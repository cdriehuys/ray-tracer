package main

import (
	"math"
	"testing"
)

func TestMakeCamera(t *testing.T) {
	width := 160
	height := 120
	fov := math.Pi / 2

	camera := MakeCamera(width, height, fov)

	if camera.Width != width {
		t.Errorf("Expected camera width to be %d, got %d", width, camera.Width)
	}

	if camera.Height != height {
		t.Errorf("Expected camera height to be %d, got %d", height, camera.Height)
	}

	if !Float64Equal(camera.FieldOfView, fov) {
		t.Errorf("Expected camera fov to be %f, got %f", fov, camera.FieldOfView)
	}

	if !IdentityMatrix4.Equals(camera.Transform) {
		t.Errorf("Expected camera transform to be %v, got %v", IdentityMatrix4, camera.Transform)
	}
}

func TestMakeCamera_PixelSize(t *testing.T) {
	testCases := []struct {
		name   string
		camera Camera
		want   float64
	}{
		{
			"horizontal canvas",
			MakeCamera(200, 125, math.Pi/2),
			0.01,
		},
		{
			"vertical canvas",
			MakeCamera(125, 200, math.Pi/2),
			0.01,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.camera.pixelSize; !Float64Equal(tt.want, got) {
				t.Errorf("Expected pixel size %f, got %f", tt.want, got)
			}
		})
	}
}

func TestCamera_MakeRayForPixel(t *testing.T) {
	testCases := []struct {
		name   string
		camera Camera
		x      int
		y      int
		want   Ray
	}{
		{
			"center of canvas",
			MakeCamera(201, 101, math.Pi/2),
			100,
			50,
			MakeRay(MakePoint(0, 0, 0), MakeVector(0, 0, -1)),
		},
		{
			"corner of canvas",
			MakeCamera(201, 101, math.Pi/2),
			0,
			0,
			MakeRay(
				MakePoint(0, 0, 0),
				MakeVector(0.66519, 0.33259, -0.66851),
			),
		},
		{
			"transformed camera",
			func() Camera {
				camera := MakeCamera(201, 101, math.Pi/2)
				camera.Transform = MakeYRotation(math.Pi / 4).
					Multiply(MakeTranslation(0, -2, 5))

				return camera
			}(),
			100,
			50,
			MakeRay(
				MakePoint(0, 2, -5),
				MakeVector(math.Sqrt2/2, 0, -math.Sqrt2/2),
			),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.camera.MakeRayForPixel(tt.x, tt.y)

			if !tt.want.Direction.Equals(got.Direction) {
				t.Errorf("Expected ray direction %v, got %v", tt.want.Direction, got.Direction)
			}

			if !tt.want.Origin.Equals(got.Origin) {
				t.Errorf("Expected ray origin %v, got %v", tt.want.Origin, got.Origin)
			}
		})
	}
}
