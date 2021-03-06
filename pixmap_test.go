package main

import (
	"strings"
	"testing"
)

func TestWriteCanvasToPPM(t *testing.T) {
	var writer strings.Builder
	source := MakeCanvas(5, 10)
	source.SetPixel(2, 4, MakeColor(1, 0, 0))
	source.SetPixel(4, 2, MakeColor(0, 1, 0))
	source.SetPixel(1, 8, MakeColor(0, 0, 1))

	var expectedWriter strings.Builder

	if err := writePPMHeader(source, &expectedWriter); err != nil {
		t.Errorf("writePPMHeader() returned error: %v", err)
	}

	if err := writePPMBody(source, &expectedWriter); err != nil {
		t.Errorf("writePPMBody() returned error: %v", err)
	}

	expected := expectedWriter.String()

	err := WriteCanvasToPPM(source, &writer)

	if err != nil {
		t.Errorf("WriteCanvasToPPM() returned error: %v", err)
	}

	if got := writer.String(); got != expected {
		t.Errorf("Expected contents to be:\n%v\nGot:\n%v", expected, got)
	}
}

func TestWritePPMBody(t *testing.T) {
	var writer strings.Builder
	source := MakeCanvas(5, 3)
	source.SetPixel(0, 0, MakeColor(1.5, 0, 0))
	source.SetPixel(2, 1, MakeColor(0, .5, 0))
	source.SetPixel(4, 2, MakeColor(-.5, 0, 1))

	// Note the trailing newline in the expected value.
	expected := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

	err := writePPMBody(source, &writer)

	if err != nil {
		t.Errorf("writePPMBody() returned error: %v", err)
	}

	if got := writer.String(); got != expected {
		t.Errorf("Expected body to be %v, got %v", expected, got)
	}
}

// No line should exceed 70 characters
func TestWritePPMBody_LongLines(t *testing.T) {
	var writer strings.Builder

	color := MakeColor(1, .8, .6)
	source := MakeCanvas(10, 2)
	for x := 0; x < source.Width; x++ {
		for y := 0; y < source.Height; y++ {
			source.SetPixel(x, y, color)
		}
	}

	expected := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`

	err := writePPMBody(source, &writer)

	if err != nil {
		t.Errorf("writePPMBody() returned error: %v", err)
	}

	if got := writer.String(); got != expected {
		t.Errorf("Expected body to be %v, got %v", expected, got)
	}
}

func TestWritePPMHeader(t *testing.T) {
	var writer strings.Builder
	source := MakeCanvas(4, 3)
	expected := "P3\n4 3\n255\n"

	err := writePPMHeader(source, &writer)

	if err != nil {
		t.Errorf("writePPMHeader() returned error: %v", err)
	}

	if got := writer.String(); got != expected {
		t.Errorf("Expected header to be %#v, got %#v", expected, got)
	}
}

func TestScaleToPPMValue(t *testing.T) {
	testCases := []struct {
		name  string
		value float64
		want  int
	}{
		{
			"zero",
			0,
			0,
		},
		{
			"one",
			1,
			255,
		},
		{
			"half",
			.5,
			128,
		},
		{
			"below range",
			-1.3,
			0,
		},
		{
			"above range",
			1.5,
			255,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := scaleToPPMValue(tt.value); got != tt.want {
				t.Errorf("Expected scaleToPPMValue(%f) = %d, got %d", tt.value, tt.want, got)
			}
		})
	}
}
