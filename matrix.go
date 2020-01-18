package main

// A square matrix, ie a 2D grid of numbers. Matrices are addressed in row-major
// order, eg matrix[1, 2] is the value at the 1st row in the 2nd column (with
// 0-based indexing).
type Matrix struct {
	Size  int
	store []float64
}

// Create a 2x2 matrix.
func MakeMatrix2(v1, v2, v3, v4 float64) Matrix {
	return Matrix{
		Size: 2,
		store: []float64{
			v1, v2,
			v3, v4,
		},
	}
}

// Create a 3x3 matrix.
func MakeMatrix3(v1, v2, v3, v4, v5, v6, v7, v8, v9 float64) Matrix {
	return Matrix{
		Size: 3,
		store: []float64{
			v1, v2, v3,
			v4, v5, v6,
			v7, v8, v9,
		},
	}
}

// Create a 4x4 matrix.
func MakeMatrix4(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16 float64) Matrix {
	return Matrix{
		Size: 4,
		store: []float64{
			v1, v2, v3, v4,
			v5, v6, v7, v8,
			v9, v10, v11, v12,
			v13, v14, v15, v16,
		},
	}
}

// Get the value at a specific row and column in the matrix.
func (m Matrix) Get(row, column int) float64 {
	return m.store[row*m.Size+column]
}

// Determine if two matrices are equal. Matrices are equal if they contain the
// same values in the same positions.
func (m Matrix) Equals(other Matrix) bool {
	if m.Size != other.Size {
		return false
	}

	for row := 0; row < m.Size; row++ {
		for col := 0; col < m.Size; col++ {
			if !Float64Equal(m.Get(row, col), other.Get(row, col)) {
				return false
			}
		}
	}

	return true
}
