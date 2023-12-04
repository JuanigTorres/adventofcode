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