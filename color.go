package main

// A representation of a color as pixel intensities. Intensities range from 0,
// all the way off, to 1, all the way on.
type Color struct {
	// The underlying data is stored in a tuple and all calculations are
	// delegated to the tuple when possible.
	tuple Tuple
}

// Get the result of adding a color to this color.
func (c Color) Add(other Color) Color {
	return Color{c.tuple.Add(other.tuple)}
}

// Get the color that is result of blending this color with another color. This
// is also known as the Hadamard product or the Schur product.
func (c Color) Blend(other Color) Color {
	return MakeColor(
		c.Red()*other.Red(),
		c.Green()*other.Green(),
		c.Blue()*other.Blue(),
	)
}

// Get the blue intensity of the color.
func (c Color) Blue() float64 {
	return c.tuple.Z
}

// Determine if another color is equivalent to this color.
func (c Color) Equals(other Color) bool {
	return c.tuple.Equals(other.tuple)
}

// Get the green intensity of the color.
func (c Color) Green() float64 {
	return c.tuple.Y
}

// Multiply the color by a scalar factor.
func (c Color) Multiply(factor float64) Color {
	return Color{c.tuple.Multiply(factor)}
}

// Get the red intensity of the color.
func (c Color) Red() float64 {
	return c.tuple.X
}

// Get the result of subtracting a color from this color
func (c Color) Subtract(other Color) Color {
	return Color{c.tuple.Subtract(other.tuple)}
}

// Create a new color object with the specified intensities of red, green, and
// blue.
func MakeColor(red, green, blue float64) Color {
	return Color{
		Tuple{
			X: red,
			Y: green,
			Z: blue,
		},
	}
}
