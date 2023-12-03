package slices

func Reduce[I, T any](arr []T, reduce func(prev I, curr T) I, initial I) I {
	result := initial
	for _, v := range arr {
		result = reduce(result, v)
	}
	return result
}