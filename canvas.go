package main

// A canvas is a 2D array of pixels. A canvas' X-coordinates lie in the range
// [0, width) and its Y-coordinates lie in the range [0, height).
type Canvas struct {
	Width  int
	Height int

	// The 2D array backing the canvas.
	colors [][]Color
}

// Get the color of a specific pixel on the canvas.
func (c Canvas) GetPixel(x, y int) Color {
	return c.colors[x][y]
}

// Set the color of a specific pixel on the canvas.
func (c *Canvas) SetPixel(x, y int, color Color) {
	c.colors[x][y] = color
}

// Create a new canvas with the specified dimensions.
func MakeCanvas(width, height int) Canvas {
	// We store colors in column-major order so that we can access normally with
	// (x, y)-like syntax.
	columns := make([][]Color, width)
	for x := range columns {
		columns[x] = make([]Color, height)
	}

	return Canvas{
		width,
		height,
		columns,
	}
}
