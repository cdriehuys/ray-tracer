package main

import "testing"

type colorOpTestCase struct {
	name  string
	base  Color
	other Color
	want  Color
}

func TestColor_Add(t *testing.T) {
	testCases := []colorOpTestCase{
		{
			"basic addition",
			MakeColor(.9, .6, .75),
			MakeColor(.7, .1, .25),
			MakeColor(1.6, .7, 1),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.base.Add(tt.other); !got.Equals(tt.want) {
				t.Errorf("Expected %v + %v = %v, got %v", tt.base, tt.other, tt.want, got)
			}
		})
	}
}

func TestColor_Blend(t *testing.T) {
	testCases := []colorOpTestCase{
		{
			"basic blend",
			MakeColor(1, .2, .4),
			MakeColor(.9, 1, .1),
			MakeColor(.9, .2, .04),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.base.Blend(tt.other); !got.Equals(tt.want) {
				t.Errorf("Expected %v * %v = %v, got %v", tt.base, tt.other, tt.want, got)
			}
		})
	}
}

func TestColor_Equals(t *testing.T) {
	testCases := []struct {
		name    string
		base    Color
		compare Color
		want    bool
	}{
		{
			"equal colors",
			MakeColor(-.5, .4, 1.7),
			MakeColor(-.5, .4, 1.7),
			true,
		},
		{
			"different red",
			MakeColor(1, 0, 0),
			MakeColor(0, 0, 0),
			false,
		},
		{
			"different green",
			MakeColor(0, 1, 0),
			MakeColor(0, 0, 0),
			false,
		},
		{
			"different blue",
			MakeColor(0, 0, 1),
			MakeColor(0, 0, 0),
			false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.base.Equals(tt.compare); got != tt.want {
				t.Errorf("Expected %v == %v = %v, got %v", tt.base, tt.compare, tt.want, got)
			}
		})
	}
}

func TestColor_Multiply(t *testing.T) {
	testCases := []struct {
		name   string
		base   Color
		factor float64
		want   Color
	}{
		{
			"multiply by scalar",
			MakeColor(.2, .3, .4),
			2,
			MakeColor(.4, .6, .8),
		},
	}
	for _, tt := range testCases {
		if got := tt.base.Multiply(tt.factor); !got.Equals(tt.want) {
			t.Errorf("Expected %v * %f = %v, got %v", tt.base, tt.factor, tt.want, got)
		}
	}
}

func TestColor_Subtract(t *testing.T) {
	testCases := []colorOpTestCase{
		{
			"basic subtraction",
			MakeColor(.9, .6, .75),
			MakeColor(.7, .1, .25),
			MakeColor(.2, .5, .5),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.base.Subtract(tt.other); !got.Equals(tt.want) {
				t.Errorf("Expected %v - %v = %v, got %v", tt.base, tt.other, tt.want, got)
			}
		})
	}
}

func TestMakeColor(t *testing.T) {
	r := -.5
	g := .4
	b := 1.7

	color := MakeColor(r, g, b)

	if !Float64Equal(color.Red(), r) {
		t.Errorf("Expected color.Red() = %f, got %f", r, color.Red())
	}

	if !Float64Equal(color.Green(), g) {
		t.Errorf("Expected color.Green() = %f, got %f", g, color.Green())
	}

	if !Float64Equal(color.Blue(), b) {
		t.Errorf("Expected color.Blue() = %f, got %f", b, color.Blue())
	}
}
