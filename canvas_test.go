package main

import "testing"

func TestCanvas_SetPixel_And_GetPixel(t *testing.T) {
	red := MakeColor(1, 0, 0)

	testCases := []struct {
		name   string
		canvas Canvas
		x      int
		y      int
	}{
		{
			"top left",
			MakeCanvas(10, 20),
			0,
			0,
		},
		{
			"top right",
			MakeCanvas(10, 20),
			9,
			0,
		},
		{
			"bottom left",
			MakeCanvas(10, 20),
			0,
			19,
		},
		{
			"bottom right",
			MakeCanvas(10, 20),
			9,
			19,
		},
		{
			"middle-ish",
			MakeCanvas(10, 20),
			4,
			3,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.canvas.SetPixel(tt.x, tt.y, red)
			if got := tt.canvas.GetPixel(tt.x, tt.y); !got.Equals(red) {
				t.Errorf("Expected canvas(%d, %d) = %v, got %v", tt.x, tt.y, red, got)
			}
		})
	}
}

func TestMakeCanvas(t *testing.T) {
	width := 10
	height := 20
	black := MakeColor(0, 0, 0)

	canvas := MakeCanvas(width, height)

	if got := canvas.Width; got != width {
		t.Errorf("Expected canvas.Width = %v, got %v", width, got)
	}

	if got := canvas.Height; got != height {
		t.Errorf("Expected canvas.Height = %v, got %v", height, got)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if got := canvas.GetPixel(x, y); !got.Equals(black) {
				t.Errorf("Expected canvas(%d, %d) = %v, got %v", x, y, black, got)
			}
		}
	}
}
