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

// Write the contents of a canvas in PPM format.
func WriteCanvasToPPM(canvas Canvas, dest io.Writer) error {
	if err := writePPMHeader(canvas, dest); err != nil {
		return fmt.Errorf("failed to write PPM header: %w", err)
	}

	if err := writePPMBody(canvas, dest); err != nil {
		return fmt.Errorf("failed to write PPM body: %w", err)
	}

	return nil
}

// Write the individual pixel data to a destination. Each pixel from the source
// canvas is scaled so that the RGB values are integers in the PPM's color range
// rather than floats in the range [0, 1].
func writePPMBody(source Canvas, dest io.Writer) error {
	for y := 0; y < source.Height; y++ {
		lineLength := 0

		for x := 0; x < source.Width; x++ {
			color := source.GetPixel(x, y)

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
					_, err := dest.Write([]byte("\n"))
					if err != nil {
						return err
					}

					lineLength = 0
				}

				// If there is already a value on the line, we want to write a
				// separator before the value.
				if lineLength != 0 {
					_, err := dest.Write([]byte(" "))
					if err != nil {
						return err
					}

					lineLength += 1
				}

				_, err := dest.Write([]byte(valueString))
				if err != nil {
					return err
				}

				lineLength += len(valueString)
			}
		}

		_, err := dest.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

// Write the header of the PPM file. The header includes the PPM version string,
// the dimensions of the image, and the maximum color value.
func writePPMHeader(canvas Canvas, dest io.Writer) error {
	contents := PPMVersion +
		"\n" +
		strconv.Itoa(canvas.Width) +
		" " +
		strconv.Itoa(canvas.Height) +
		"\n" +
		strconv.Itoa(PPMMaxColorValue) +
		"\n"

	_, err := dest.Write([]byte(contents))

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
