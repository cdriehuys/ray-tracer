package main

import (
	"testing"
)

func TestViewTransform(t *testing.T) {
	tests := []struct {
		name string
		from Tuple
		to   Tuple
		up   Tuple
		want Matrix
	}{
		{
			"default orientation",
			MakeVector(0, 0, 0),
			MakeVector(0, 0, -1),
			MakeVector(0, 1, 0),
			IdentityMatrix4,
		},
		{
			"looking in positive z direction",
			MakeVector(0, 0, 0),
			MakeVector(0, 0, 1),
			MakeVector(0, 1, 0),
			MakeScale(-1, 1, -1),
		},
		{
			"transform",
			MakePoint(0, 0, 8),
			MakePoint(0, 0, 0),
			MakeVector(0, 1, 0),
			MakeTranslation(0, 0, -8),
		},
		{
			"arbitrary transform",
			MakePoint(1, 3, 2),
			MakePoint(4, -2, 8),
			MakeVector(1, 1, 0),
			MakeMatrix4(
				-0.50709, 0.50709, 0.67612, -2.36643,
				0.76772, 0.60609, 0.12122, -2.82843,
				-0.35857, 0.59761, -0.71714, 0.00000,
				0.00000, 0.00000, 0.00000, 1.00000,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ViewTransform(tt.from, tt.to, tt.up); !tt.want.Equals(got) {
				t.Errorf("ViewTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
