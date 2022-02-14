package main

import "fmt"

type Resulter[V any] interface {
	Value() V
	Error() error
}

type Function[T any, V any] struct {
	arg    T
	action func(T) Resulter[V]
}

func (f *Function[T, V]) Do() Resulter[V] {
return f.action(f.arg)
}

func NewFunc[T any, V any](f func(T) Resulter[V], arg T) Function[T, V] {
	return Function[T, V]{
		arg:    arg,
		action: f,
	}
}

type Result[V any] struct {
	val V
	err error
}

func (r *Result[V]) Value() V {
	return r.val
}

func (r *Result[V]) Error() error {
	return r.err
}

func main() {
	f := NewFunc(func(val bool) Resulter[int32] {
		if val {
			return &Result[int32]{
				val: int32(1),
				err: nil,
			}
		}
		return &Result[int32]{
			val: int32(0),
			err: nil,
		}
	}, true)

	fmt.Println(f.Do())
}
