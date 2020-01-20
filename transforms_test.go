package main

import (
	"math"
	"testing"
)

type transformTestCase struct {
	name      string
	source    Tuple
	transform Matrix
	want      Tuple
}

func TestMakeRotation(t *testing.T) {
	testCases := []transformTestCase{
		{
			"x rotation half quarter",
			MakePoint(0, 1, 0),
			MakeXRotation(math.Pi / 4),
			MakePoint(0, math.Sqrt2/2, math.Sqrt2/2),
		},
		{
			"x rotation half quarter inverse",
			MakePoint(0, 1, 0),
			MakeXRotation(math.Pi / 4).Inverted(),
			MakePoint(0, math.Sqrt2/2, -math.Sqrt2/2),
		},
		{
			"x rotation full quarter",
			MakePoint(0, 1, 0),
			MakeXRotation(math.Pi / 2),
			MakePoint(0, 0, 1),
		},
		{
			"y rotation half quarter",
			MakePoint(0, 0, 1),
			MakeYRotation(math.Pi / 4),
			MakePoint(math.Sqrt2/2, 0, math.Sqrt2/2),
		},
		{
			"y rotation half quarter inverse",
			MakePoint(0, 0, 1),
			MakeYRotation(math.Pi / 4).Inverted(),
			MakePoint(-math.Sqrt2/2, 0, math.Sqrt2/2),
		},
		{
			"y rotation full quarter",
			MakePoint(0, 0, 1),
			MakeYRotation(math.Pi / 2),
			MakePoint(1, 0, 0),
		},
		{
			"z rotation half quarter",
			MakePoint(0, 1, 0),
			MakeZRotation(math.Pi / 4),
			MakePoint(-math.Sqrt2/2, math.Sqrt2/2, 0),
		},
		{
			"z rotation half quarter inverse",
			MakePoint(0, 1, 0),
			MakeZRotation(math.Pi / 4).Inverted(),
			MakePoint(math.Sqrt2/2, math.Sqrt2/2, 0),
		},
		{
			"z rotation full quarter",
			MakePoint(0, 1, 0),
			MakeZRotation(math.Pi / 2),
			MakePoint(-1, 0, 0),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.transform.TupleMultiply(tt.source); !got.Equals(tt.want) {
				t.Errorf("Rotation expectation failed:\nRotation: %v, \nSource: %v\nExpected: %v\nReceived: %v", tt.transform, tt.source, tt.want, got)
			}
		})
	}
}

func TestMakeScale(t *testing.T) {
	testCases := []transformTestCase{
		{
			"point",
			MakePoint(-4, 6, 8),
			MakeScale(2, 3, 4),
			MakePoint(-8, 18, 32),
		},
		{
			"point inverse",
			MakePoint(-4, 6, 8),
			MakeScale(2, 3, 4).Inverted(),
			MakePoint(-2, 2, 2),
		},
		{
			"vector",
			MakeVector(-4, 6, 8),
			MakeScale(2, 3, 4),
			MakeVector(-8, 18, 32),
		},
		{
			"x-axis reflection",
			MakePoint(2, 3, 4),
			MakeScale(-1, 1, 1),
			MakePoint(-2, 3, 4),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.transform.TupleMultiply(tt.source); !got.Equals(tt.want) {
				t.Errorf("Scale expectation failed:\nScale: %v, \nSource: %v\nExpected: %v\nReceived: %v", tt.transform, tt.source, tt.want, got)
			}
		})
	}
}

func TestMakeShear(t *testing.T) {
	testCases := []transformTestCase{
		{
			"x in proportion to y",
			MakePoint(2, 3, 4),
			MakeShear(1, 0, 0, 0, 0, 0),
			MakePoint(5, 3, 4),
		},
		{
			"x in proportion to z",
			MakePoint(2, 3, 4),
			MakeShear(0, 1, 0, 0, 0, 0),
			MakePoint(6, 3, 4),
		},
		{
			"y in proportion to x",
			MakePoint(2, 3, 4),
			MakeShear(0, 0, 1, 0, 0, 0),
			MakePoint(2, 5, 4),
		},
		{
			"y in proportion to z",
			MakePoint(2, 3, 4),
			MakeShear(0, 0, 0, 1, 0, 0),
			MakePoint(2, 7, 4),
		},
		{
			"z in proportion to x",
			MakePoint(2, 3, 4),
			MakeShear(0, 0, 0, 0, 1, 0),
			MakePoint(2, 3, 6),
		},
		{
			"z in proportion to y",
			MakePoint(2, 3, 4),
			MakeShear(0, 0, 0, 0, 0, 1),
			MakePoint(2, 3, 7),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.transform.TupleMultiply(tt.source); !got.Equals(tt.want) {
				t.Errorf("Shear expectation failed:\nShear: %v, \nSource: %v\nExpected: %v\nReceived: %v", tt.transform, tt.source, tt.want, got)
			}
		})
	}
}

func TestMakeTranslation(t *testing.T) {
	testCases := []transformTestCase{
		{
			"point",
			MakePoint(-3, 4, 5),
			MakeTranslation(5, -3, 2),
			MakePoint(2, 1, 7),
		},
		{
			"point inverse",
			MakePoint(-3, 4, 5),
			MakeTranslation(5, -3, 2).Inverted(),
			MakePoint(-8, 7, 3),
		},
		{
			"vector",
			MakeVector(-3, 4, 5),
			MakeTranslation(5, -3, 2),
			MakeVector(-3, 4, 5),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.transform.TupleMultiply(tt.source); !got.Equals(tt.want) {
				t.Errorf("Translation expectation failed:\nTranslation: %v, \nSource: %v\nExpected: %v\nReceived: %v", tt.transform, tt.source, tt.want, got)
			}
		})
	}
}

func TestTransformationCombination(t *testing.T) {
	p := MakePoint(1, 0, 1)
	A := MakeXRotation(math.Pi / 2)
	B := MakeScale(5, 5, 5)
	C := MakeTranslation(10, 5, 7)

	p2 := A.TupleMultiply(p)
	p2Expected := MakePoint(1, -1, 0)
	if !p2.Equals(p2Expected) {
		t.Errorf("Expected %v\nReceived %v", p2Expected, p2)
	}

	p3 := B.TupleMultiply(p2)
	p3Expected := MakePoint(5, -5, 0)
	if !p3.Equals(p3Expected) {
		t.Errorf("Expected %v\nReceived %v", p3Expected, p3)
	}

	p4 := C.TupleMultiply(p3)
	p4Expected := MakePoint(15, 0, 7)
	if !p4.Equals(p4Expected) {
		t.Errorf("Expected %v\nReceived %v", p4Expected, p4)
	}

	// The same result should be achievable by chaining the transformations.
	p5 := C.Multiply(B).Multiply(A).TupleMultiply(p)
	p5Expected := MakePoint(15, 0, 7)
	if !p5.Equals(p5Expected) {
		t.Errorf("Expected %v\nReceived %v", p5Expected, p5)
	}
}
