package test

type typ[T any] struct {
	t T
}

type intLike interface {
	uint8 | int
}

func max[T intLike](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func max3[T intLike](a, b, c T) T {
	return max[T](max[T](a, b), c)
}

func (typ[T]) f() {}

func (*typ[T]) f2() {}
