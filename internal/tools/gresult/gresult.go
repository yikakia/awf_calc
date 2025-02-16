package gresult

import (
	"fmt"
)

type Result[T any] struct {
	val T
	err error
}

func Of[T any](val T, err error) Result[T] {
	return Result[T]{val: val, err: err}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return !r.IsOk()
}

func (r Result[T]) Get() (T, error) {
	return r.val, r.err
}

func (r Result[T]) MustGet() T {
	if !r.IsOk() {
		panic(fmt.Sprintf("err should be nil.err:%v", r.err))
	}
	return r.val
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) ValueOr(fallback T) T {
	if r.IsErr() {
		return fallback
	}
	return r.val
}
func (r Result[T]) ValueOrZero() (t T) {
	if r.IsErr() {
		return t
	}
	return r.val
}
