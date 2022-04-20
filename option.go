package opt

import "encoding/json"

type Option[T any] struct {
	val *T
}

func New[T any](val T) Option[T] {
	return Option[T]{val: &val}
}

func (o Option[T]) Some() (T, bool) {
	if o.val == nil {
		return *new(T), false
	}
	return *o.val, true
}

func (o Option[T]) None() bool {
	return o.val == nil
}

func (o Option[T]) Unwrap() T {
	return *o.val
}

func (o Option[T]) UnwrapOr(val T) T {
	if o.None() {
		return val
	}
	return *o.val
}

func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.None() {
		return f()
	}
	return *o.val
}

func (o Option[T]) UnwrapOrDefault() T {
	if o.val == nil {
		return *new(T)
	}
	return *o.val
}

func (o *Option[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.val)
}

func (o *Option[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &o.val)
}
