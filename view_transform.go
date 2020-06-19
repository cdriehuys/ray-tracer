package main

// Create a view transformation representing the case where the observer is at
// the point `from`, looking towards the point `to`, with the vector `up`
// pointing in the upward direction.
func ViewTransform(from, to, up Tuple) Matrix {
	forward := to.Subtract(from).Normalized()
	left := forward.Cross(up.Normalized())
	// We recalculate up so that the provided `up` does not have to be precisely
	// computed.
	trueUp := left.Cross(forward)

	orientation := MakeMatrix4(
		left.X, left.Y, left.Z, 0,
		trueUp.X, trueUp.Y, trueUp.Z, 0,
		-forward.X, -forward.Y, -forward.Z, 0,
		0, 0, 0, 1,
	)

	return orientation.Multiply(MakeTranslation(-from.X, -from.Y, -from.Z))
}
