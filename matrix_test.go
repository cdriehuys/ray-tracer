package main

import (
	"fmt"
	"testing"
)

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

func TestMatrix_Cofactor(t *testing.T) {
	type matrixTestCase struct {
		row          int
		col          int
		wantMinor    float64
		wantCofactor float64
	}
	testCases := []struct {
		name   string
		matrix Matrix
		tests  []matrixTestCase
	}{
		{
			"3x3",
			MakeMatrix3(
				3, 5, 0,
				2, -1, -7,
				6, -1, 5,
			),
			[]matrixTestCase{
				{0, 0, -12, -12},
				{1, 0, 25, -25},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			for _, testCase := range tt.tests {
				t.Run(fmt.Sprintf("Row %d Column %d", testCase.row, testCase.col), func(t *testing.T) {
					minor := tt.matrix.Minor(testCase.row, testCase.col)
					if !Float64Equal(minor, testCase.wantMinor) {
						t.Errorf("Expected minor(A, %d, %d) = %v, got %v\nA = %v", testCase.row, testCase.col, testCase.wantMinor, minor, tt.matrix)
					}

					cofactor := tt.matrix.Cofactor(testCase.row, testCase.col)
					if !Float64Equal(cofactor, testCase.wantCofactor) {
						t.Errorf("Expected cofactor(A, %d, %d) = %v, got %v\nA = %v", testCase.row, testCase.col, testCase.wantCofactor, cofactor, tt.matrix)
					}
				})
			}
		})
	}
}

func TestMatrix_Determinant(t *testing.T) {
	type cofactorTest struct {
		row  int
		col  int
		want float64
	}
	testCases := []struct {
		name          string
		a             Matrix
		cofactorTests []cofactorTest
		want          float64
	}{
		{
			name: "2x2",
			a: MakeMatrix2(
				1, 5,
				-3, 2,
			),
			want: 17,
		},
		{
			name: "3x3",
			a: MakeMatrix3(
				1, 2, 6,
				-5, 8, -4,
				2, 6, 4,
			),
			cofactorTests: []cofactorTest{
				{0, 0, 56},
				{0, 1, 12},
				{0, 2, -46},
			},
			want: -196,
		},
		{
			name: "4x4",
			a: MakeMatrix4(
				-2, -8, 3, 5,
				-3, 1, 7, 3,
				1, 2, -9, 6,
				-6, 7, 7, -9,
			),
			cofactorTests: []cofactorTest{
				{0, 0, 690},
				{0, 1, 447},
				{0, 2, 210},
				{0, 3, 51},
			},
			want: -4071,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			for _, ct := range tt.cofactorTests {
				t.Run(fmt.Sprintf("cofactor row %d column %d", ct.row, ct.col), func(t *testing.T) {
					if got := tt.a.Cofactor(ct.row, ct.col); !Float64Equal(got, ct.want) {
						t.Errorf("Expected cofactor(A, %d, %d) = %v, got %v\nA = %v", ct.row, ct.col, ct.want, got, tt.a)
					}
				})
			}

			if got := tt.a.Determinant(); !Float64Equal(got, tt.want) {
				t.Errorf("Expected determinant(A) = %v, got %v\nA = %v", tt.want, got, tt.a)
			}
		})
	}
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

func TestMatrix_Minor(t *testing.T) {
	testCases := []struct {
		name string
		a    Matrix
		row  int
		col  int
		want float64
	}{
		{
			"3x3",
			MakeMatrix3(
				3, 5, 0,
				2, -1, -7,
				6, -1, 5,
			),
			1,
			0,
			25,
		},
	}
	for _, tt := range testCases {
		b := tt.a.Submatrix(tt.row, tt.col)
		determinant := b.Determinant()

		if !Float64Equal(determinant, tt.want) {
			t.Errorf("Expected determinant(submatrix(A, %d, %d)) = %v, got %v\nA = %v", tt.row, tt.col, tt.want, determinant, tt.a)
		}

		minor := tt.a.Minor(tt.row, tt.col)
		if !Float64Equal(minor, tt.want) {
			t.Errorf("Expected minor(A, %d, %d) = %v, got %v\nA = %v", tt.row, tt.col, tt.want, minor, tt.a)
		}
	}
}

func TestMatrix_Multiply(t *testing.T) {
	testCases := []struct {
		name string
		a    Matrix
		b    Matrix
		want Matrix
	}{
		{
			"4x4",
			MakeMatrix4(
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			),
			MakeMatrix4(
				-2, 1, 2, 3,
				3, 2, 1, -1,
				4, 3, 6, 5,
				1, 2, 7, 8,
			),
			MakeMatrix4(
				20, 22, 50, 48,
				44, 54, 114, 108,
				40, 58, 110, 102,
				16, 26, 46, 42,
			),
		},
		{
			"4x4 identity",
			MakeMatrix4(
				0, 1, 2, 4,
				1, 2, 4, 8,
				2, 4, 8, 16,
				4, 8, 16, 32,
			),
			IdentityMatrix4,
			MakeMatrix4(
				0, 1, 2, 4,
				1, 2, 4, 8,
				2, 4, 8, 16,
				4, 8, 16, 32,
			),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Multiply(tt.b); !got.Equals(tt.want) {
				t.Errorf("Expected A x B = C, got D:\nA = %v\nB = %v\nC = %v\nD = %v", tt.a, tt.b, tt.want, got)
			}
		})
	}
}

func TestMatrix_Submatrix(t *testing.T) {
	testCases := []struct {
		name string
		a    Matrix
		row  int
		col  int
		want Matrix
	}{
		{
			"3x3 to 2x2",
			MakeMatrix3(
				1, 5, 0,
				-3, 2, 7,
				0, 6, -3,
			),
			0,
			2,
			MakeMatrix2(
				-3, 2,
				0, 6,
			),
		},
		{
			"4x4 to 3x3",
			MakeMatrix4(
				-6, 1, 1, 6,
				-8, 5, 8, 6,
				-1, 0, 8, 2,
				-7, 1, -1, 1,
			),
			2,
			1,
			MakeMatrix3(
				-6, 1, 6,
				-8, 8, 6,
				-7, -1, 1,
			),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Submatrix(tt.row, tt.col); !got.Equals(tt.want) {
				t.Errorf("Expected submatrix(A, %d, %d) = B, got C\nA = %v\nB = %v\nC = %v", tt.row, tt.col, tt.a, tt.want, got)
			}
		})
	}
}

func TestMatrix_Transposed(t *testing.T) {
	testCases := []struct {
		name string
		a    Matrix
		want Matrix
	}{
		{
			"4x4 identity",
			IdentityMatrix4,
			IdentityMatrix4,
		},
		{
			"4x4",
			MakeMatrix4(
				0, 9, 3, 0,
				9, 8, 0, 8,
				1, 8, 5, 3,
				0, 0, 5, 8,
			),
			MakeMatrix4(
				0, 9, 1, 0,
				9, 8, 8, 0,
				3, 0, 5, 5,
				0, 8, 3, 8,
			),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Transposed(); !got.Equals(tt.want) {
				t.Errorf("Expected transpose(A) = B, got C\nA = %v\nB = %v\nC = %v", tt.a, tt.want, got)
			}
		})
	}
}

func TestMatrix_TupleMultiply(t *testing.T) {
	testCases := []struct {
		name string
		a    Matrix
		b    Tuple
		want Tuple
	}{
		{
			"4x4",
			MakeMatrix4(
				1, 2, 3, 4,
				2, 4, 4, 2,
				8, 6, 4, 1,
				0, 0, 0, 1,
			),
			Tuple{1, 2, 3, 1},
			Tuple{18, 24, 33, 1},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.TupleMultiply(tt.b); !got.Equals(tt.want) {
				t.Errorf("Expected A x b = c, got d:\nA = %v\nb = %v\nc = %v\nd = %v", tt.a, tt.b, tt.want, got)
			}
		})
	}
}

func assertMatrixValue(t *testing.T, m Matrix, row, column int, expected float64) {
	if got := m.Get(row, column); !Float64Equal(got, expected) {
		t.Errorf("Expected matrix[%d, %d] = %v, got %v", row, column, expected, got)
	}
}
