package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

const (
	PPMVersion       = "P3"
	PPMMaxColorValue = 255
	PPMMaxLineLength = 70
)

type Pixmap struct {
	Source Canvas
	Dest   io.Writer
}

// Write the Pixmap's contents to its writer.
func (p *Pixmap) Write() error {
	if err := p.writeHeader(); err != nil {
		return fmt.Errorf("failed to write PPM header: %w", err)
	}

	if err := p.writeBody(); err != nil {
		return fmt.Errorf("failed to write PPM body: %w", err)
	}

	return nil
}

// Write the individual pixel data to the Pixmap's writer.
func (p *Pixmap) writeBody() error {
	for y := 0; y < p.Source.Height; y++ {
		lineLength := 0

		for x := 0; x < p.Source.Width; x++ {
			color := p.Source.GetPixel(x, y)

			for _, value := range []float64{color.Red(), color.Green(), color.Blue()} {
				valueString := strconv.Itoa(scaleToPPMValue(value))
				valueLength := len(valueString)
				if lineLength != 0 {
					// Include separator
					valueLength += 1
				}

				// If writing the value is going to exceed the max line length,
				// start a new line before writing the value.
				if lineLength+valueLength > PPMMaxLineLength {
					_, err := p.Dest.Write([]byte("\n"))
					if err != nil {
						return err
					}

					lineLength = 0
				}

				// If there is already a value on the line, we want to write a
				// separator before the value.
				if lineLength != 0 {
					_, err := p.Dest.Write([]byte(" "))
					if err != nil {
						return err
					}

					lineLength += 1
				}

				_, err := p.Dest.Write([]byte(valueString))
				if err != nil {
					return err
				}

				lineLength += len(valueString)
			}
		}

		_, err := p.Dest.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

// Write the header of the PPM file to the Pixmap's writer. The header includes
// the PPM version string, the dimensions of the image, and the maximum color
// value.
func (p *Pixmap) writeHeader() error {
	contents := PPMVersion +
		"\n" +
		strconv.Itoa(p.Source.Width) +
		" " +
		strconv.Itoa(p.Source.Height) +
		"\n" +
		strconv.Itoa(PPMMaxColorValue) +
		"\n"

	_, err := p.Dest.Write([]byte(contents))

	return err
}

// Scale a value from the range [0, 1] to a value in the range
// [0, PPM max color value]. Inputs outside the range [0, 1] are accepted but
// will be clamped to the acceptable range.
func scaleToPPMValue(value float64) int {
	if value < 0 {
		value = 0
	} else if value > 1 {
		value = 1
	}

	return int(math.Round(value * PPMMaxColorValue))
}
