package optional

type O[T any] struct {
	val T
	ok  bool
}

func New[T any](val T, ok bool) O[T] {
	return O[T]{val: val, ok: ok}
}

func Ok[T any](val T) O[T] {
	return O[T]{val: val, ok: true}
}

func None[T any]() O[T] {
	return O[T]{ok: false}
}

func (o O[T]) IsOk() bool {
	return o.ok
}

func (o O[T]) IsNil() bool {
	return !o.ok
}

func (o O[T]) Get() T {
	return o.val
}

// MustGet
// warning may panic!!
func (o O[T]) MustGet() T {
	if !o.ok {
		panic("value must be true")
	}
	return o.val
}
