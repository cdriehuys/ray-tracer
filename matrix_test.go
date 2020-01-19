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

func TestMatrix_Inverted(t *testing.T) {
	type cofactorCheck struct {
		row  int
		col  int
		want float64
	}
	testCases := []struct {
		name            string
		matrix          Matrix
		cofactorChecks  []cofactorCheck
		wantDeterminant float64
		wantInverse     Matrix
	}{
		{
			name: "4x4 #1",
			matrix: MakeMatrix4(
				-5, 2, 6, -8,
				1, -5, 1, 8,
				7, 7, -6, -7,
				1, -3, 7, 4,
			),
			cofactorChecks: []cofactorCheck{
				{2, 3, -160},
				{3, 2, 105},
			},
			wantDeterminant: 532,
			wantInverse: MakeMatrix4(
				0.21805, 0.45113, 0.24060, -0.04511,
				-0.80827, -1.45677, -0.44361, 0.52068,
				-0.07895, -0.22368, -0.05263, 0.19737,
				-0.52256, -0.81391, -0.30075, 0.30639,
			),
		},
		{
			name: "4x4 #2",
			matrix: MakeMatrix4(
				8, -5, 9, 2,
				7, 5, 6, 1,
				-6, 0, 9, 6,
				-3, 0, -9, -4,
			),
			wantInverse: MakeMatrix4(
				-0.15385, -0.15385, -0.28205, -0.53846,
				-0.07692, 0.12308, 0.02564, 0.03077,
				0.35897, 0.35897, 0.43590, 0.92308,
				-0.69231, -0.69231, -0.76923, -1.92308,
			),
		},
		{
			name: "4x4 #2",
			matrix: MakeMatrix4(
				9, 3, 0, 9,
				-5, -2, -6, -3,
				-4, 9, 6, 4,
				-7, 6, 6, 2,
			),
			wantInverse: MakeMatrix4(
				-0.04074, -0.07778, 0.14444, -0.22222,
				-0.07778, 0.03333, 0.36667, -0.33333,
				-0.02901, -0.14630, -0.10926, 0.12963,
				0.17778, 0.06667, -0.26667, 0.33333,
			),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Check the determinant if one was provided.
			determinant := tt.matrix.Determinant()
			if tt.wantDeterminant != 0 && !Float64Equal(determinant, tt.wantDeterminant) {
				t.Errorf("Expected determinant(A) = %v, got %v\nA = %v", tt.wantDeterminant, determinant, tt.matrix)
			}

			inverse := tt.matrix.Inverted()

			// If there are defined spot checks, ensure the results match the
			// expectations.
			for _, ct := range tt.cofactorChecks {
				t.Run(fmt.Sprintf("cofactor check row %d column %d", ct.row, ct.col), func(t *testing.T) {
					cofactor := tt.matrix.Cofactor(ct.row, ct.col)
					if !Float64Equal(cofactor, ct.want) {
						t.Errorf("Expected cofactor(A, %d, %d) = %v, got %v\nA = %v", ct.row, ct.col, ct.want, cofactor, tt.matrix)
					}

					expectedValue := cofactor / determinant
					if got := inverse.Get(ct.col, ct.row); !Float64Equal(got, expectedValue) {
						t.Errorf("Expected B[%d, %d] = %v, got %v\nB = %v", ct.col, ct.row, expectedValue, got, inverse)
					}
				})
			}

			// Test the expectation for the inverse as a whole.
			if !inverse.Equals(tt.wantInverse) {
				t.Errorf("Expected inverse(A) = B, got C\nA = %v\nB = %v\nC = %v", tt.matrix, tt.wantInverse, inverse)
			}
		})
	}
}

func TestMatrix_IsInvertible(t *testing.T) {
	testCases := []struct {
		name            string
		matrix          Matrix
		wantDeterminant float64
		wantInvertible  bool
	}{
		{
			"4x4 invertible",
			MakeMatrix4(
				6, 4, 4, 4,
				5, 5, 7, 6,
				4, -9, 3, -7,
				9, 1, 7, -6,
			),
			-2120,
			true,
		},
		{
			"4x4 not invertible",
			MakeMatrix4(
				-4, 2, -2, -3,
				9, 6, 2, 6,
				0, -5, 1, -5,
				0, 0, 0, 0,
			),
			0,
			false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Determinant(); !Float64Equal(got, tt.wantDeterminant) {
				t.Errorf("Expected determinant(A) = %v, got %v\nA = %v", tt.wantDeterminant, got, tt.matrix)
			}

			if got := tt.matrix.IsInvertible(); got != tt.wantInvertible {
				t.Errorf("Expected isInvertible(A) = %v, got %v\nA = %v", tt.wantInvertible, got, tt.matrix)
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

func TestMatrix_Multiply_Inverse(t *testing.T) {
	a := MakeMatrix4(
		3, -9, 7, 3,
		3, -8, 2, -9,
		-4, 4, 4, 1,
		-6, 5, -1, 1,
	)
	b := MakeMatrix4(
		8, 2, 2, 2,
		3, -1, 7, 0,
		7, 0, 5, 4,
		6, -2, 0, 5,
	)
	c := a.Multiply(b)

	if got := c.Multiply(b.Inverted()); !got.Equals(a) {
		t.Errorf("Expected A * B * inverse(B) = A, got C\nA = %v\nB = %v\nC = %v", a, b, got)
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
