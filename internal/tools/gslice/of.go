package gslice

func Of[T any](elems ...T) (r []T) {
	r = make([]T, len(elems))
	copy(r, elems)
	return r
}
