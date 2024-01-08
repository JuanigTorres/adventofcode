package functional

type (
	UnaryFunc[T, R any] func (T) R
	BinaryFunc[T1, T2, R any] func (T1, T2) R
)

func Pipe1[T1, R any](f func (T1) R) UnaryFunc[T1, R] {
	return func(x T1) R {
		return f(x)
	}
}

func Pipe4[T1, T2, T3, T4, R any](f1 UnaryFunc[T1, T2], f2 UnaryFunc[T2, T3], f3 UnaryFunc[T3, T4], f4 UnaryFunc[T4, R]) UnaryFunc[T1, R] {
	return func(x T1) R {
		return f4(f3(f2(f1(x))))
	}
}

type (
	Curried1[T1, R any] func(T1) R
	Curried2[T1, T2, R any] func(T1) func(T2) R
)

func Curry2[T1, T2, R any](f func (T1, T2) R) Curried2[T1, T2, R] {
	return func(t1 T1) func(T2) R {
		return func(t2 T2) R {
			return f(t1, t2)
		} 
	}
}

func InverseCurry2[T1, T2, R any](f func (T1, T2) R) Curried2[T2, T1, R] {
	return func(t2 T2) func(T1) R {
		return func(t1 T1) R {
			return f(t1, t2)
		} 
	}
}