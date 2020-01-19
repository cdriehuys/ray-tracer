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

var (
	IdentityMatrix4 = MakeMatrix4(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
)

// Find a cofactor in the matrix.
func (m Matrix) Cofactor(row, col int) float64 {
	minor := m.Minor(row, col)

	if (row+col)%2 == 1 {
		return -minor
	}

	return minor
}

// Find the determinant of the matrix.
func (m Matrix) Determinant() float64 {
	if m.Size == 2 {
		return m.Get(0, 0)*m.Get(1, 1) - m.Get(0, 1)*m.Get(1, 0)
	}

	determinant := 0.0
	for col := 0; col < m.Size; col++ {
		determinant += m.Get(0, col) * m.Cofactor(0, col)
	}

	return determinant
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

// Find the minor of a matrix which is defined as the determinant of its
// submatrix. The submatrix is determined by the given row and column, which are
// removed from the matrix to produce the submatrix.
func (m Matrix) Minor(row, col int) float64 {
	return m.Submatrix(row, col).Determinant()
}

// Multiply this matrix by another.
func (m Matrix) Multiply(b Matrix) Matrix {
	c := Matrix{Size: m.Size, store: make([]float64, m.Size*m.Size)}

	for row := 0; row < m.Size; row++ {
		for col := 0; col < m.Size; col++ {
			cellValue := 0.0
			for i := 0; i < m.Size; i++ {
				aVal := m.Get(row, i)
				bVal := b.Get(i, col)
				cellValue += aVal * bVal
			}
			c.Set(row, col, cellValue)
		}
	}

	return c
}

func (m *Matrix) Set(row, column int, value float64) {
	m.store[row*m.Size+column] = value
}

// Create a matrix that is a copy of the current matrix but with a specific row
// and column removed.
func (m Matrix) Submatrix(removedRow, removedCol int) Matrix {
	targetRow := 0
	targetCol := 0
	c := Matrix{
		Size:  m.Size - 1,
		store: make([]float64, (m.Size-1)*(m.Size-1)),
	}

	for row := 0; row < m.Size; row++ {
		if row == removedRow {
			continue
		}

		targetCol = 0
		for col := 0; col < m.Size; col++ {
			if col == removedCol {
				continue
			}

			c.Set(targetRow, targetCol, m.Get(row, col))

			targetCol++
		}
		targetRow++
	}

	return c
}

// Get the transpose of the target matrix.
func (m Matrix) Transposed() Matrix {
	c := Matrix{
		Size:  m.Size,
		store: make([]float64, m.Size*m.Size),
	}

	for row := 0; row < m.Size; row++ {
		for col := 0; col < m.Size; col++ {
			c.Set(col, row, m.Get(row, col))
		}
	}

	return c
}

// Multiply the target matrix by a tuple.
func (m Matrix) TupleMultiply(b Tuple) Tuple {
	return Tuple{
		X: m.Get(0, 0)*b.X +
			m.Get(0, 1)*b.Y +
			m.Get(0, 2)*b.Z +
			m.Get(0, 3)*b.W,
		Y: m.Get(1, 0)*b.X +
			m.Get(1, 1)*b.Y +
			m.Get(1, 2)*b.Z +
			m.Get(1, 3)*b.W,
		Z: m.Get(2, 0)*b.X +
			m.Get(2, 1)*b.Y +
			m.Get(2, 2)*b.Z +
			m.Get(2, 3)*b.W,
		W: m.Get(3, 0)*b.X +
			m.Get(3, 1)*b.Y +
			m.Get(3, 2)*b.Z +
			m.Get(3, 3)*b.W,
	}
}
