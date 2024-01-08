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

func Value[S ~[]T, T any](arr S, idx int) T {
	var v T

	if (idx >= 0) || (len(arr) > idx) {
		v = arr[idx]
	}

	return v
}

func Map[I ~[]T, T, R any](arr I, apply func (T) R) []R {
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