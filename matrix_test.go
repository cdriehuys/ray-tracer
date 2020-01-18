package main

import "testing"

func TestMakeMatrix2(t *testing.T) {
	matrix := MakeMatrix2(
		-3, 5,
		1, -2,
	)

	assertMatrixValue(t, matrix, 0, 0, -3)
	assertMatrixValue(t, matrix, 0, 1, 5)
	assertMatrixValue(t, matrix, 1, 0, 1)
	assertMatrixValue(t, matrix, 1, 1, -2)
}

func TestMakeMatrix3(t *testing.T) {
	matrix := MakeMatrix3(
		-3, 5, 0,
		1, -2, -7,
		0, 1, 1,
	)

	assertMatrixValue(t, matrix, 0, 0, -3)
	assertMatrixValue(t, matrix, 1, 1, -2)
	assertMatrixValue(t, matrix, 2, 2, 1)
}

func TestMakeMatrix4(t *testing.T) {
	matrix := MakeMatrix4(
		1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5, 16.5,
	)

	assertMatrixValue(t, matrix, 0, 0, 1)
	assertMatrixValue(t, matrix, 0, 3, 4)
	assertMatrixValue(t, matrix, 1, 0, 5.5)
	assertMatrixValue(t, matrix, 1, 2, 7.5)
	assertMatrixValue(t, matrix, 2, 2, 11)
	assertMatrixValue(t, matrix, 3, 0, 13.5)
	assertMatrixValue(t, matrix, 3, 2, 15.5)
}

func TestMatrix_Equals(t *testing.T) {
	testCases := []struct {
		name string
		m1   Matrix
		m2   Matrix
		want bool
	}{
		{
			"equal 4x4",
			MakeMatrix4(
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 10, 11, 12,
				13, 14, 15, 16,
			),
			MakeMatrix4(
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 10, 11, 12,
				13, 14, 15, 16,
			),
			true,
		}, {
			"unequal 4x4",
			MakeMatrix4(
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			),
			MakeMatrix4(
				2, 3, 4, 5,
				6, 7, 8, 9,
				8, 7, 6, 5,
				4, 3, 2, 1,
			),
			false,
		},
		{
			"unequal sizes",
			MakeMatrix2(
				1, 2,
				3, 4,
			),
			MakeMatrix3(
				1, 2, 0,
				3, 4, 0,
				0, 0, 0,
			),
			false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m1.Equals(tt.m2); got != tt.want {
				if tt.want {
					t.Errorf("Expected A = B\nA = %v\nB = %v", tt.m1, tt.m2)
				} else {
					t.Errorf("Expected A != B\nA = %v\nB = %v", tt.m1, tt.m2)

				}
			}
		})
	}
}

func assertMatrixValue(t *testing.T, m Matrix, row, column int, expected float64) {
	if got := m.Get(row, column); !Float64Equal(got, expected) {
		t.Errorf("Expected matrix[%d, %d] = %v, got %v", row, column, expected, got)
	}
}
