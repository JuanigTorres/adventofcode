package slices

func Reduce[I, T any](arr []T, reduce func(prev I, curr T) I, initial I) I {
	result := initial
	for _, v := range arr {
		result = reduce(result, v)
	}
	return result
}

func Some[T any](arr []T, apply func(t T) bool) bool {
	for _, v := range arr {
		if apply(v) {
			return true
		}
	}
	return false
}

func Map[T, R any](arr []T, apply func (T) R) []R {
	r := make([]R, 0, len(arr))

	for _, a := range arr {
		r = append(r, apply(a))
	}

	return r
}

func Filter[T any](arr []T, apply func (T) bool) []T {
	r := make([]T, 0, len(arr))

	for _, a := range arr {
		if !apply(a){
			r = append(r, a)
		}
	}

	return r
}