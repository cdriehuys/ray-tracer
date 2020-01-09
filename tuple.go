package main

type Tuple struct {
	X float64
	Y float64
	Z float64
	W int8
}

func (t *Tuple) IsPoint() bool {
	return t.W == 1
}

func (t *Tuple) IsVector() bool {
	return t.W == 0
}

func MakePoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func MakeVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}
