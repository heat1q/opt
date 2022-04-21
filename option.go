package opt

import "encoding/json"

// Option represent an optional value.
type Option[T any] struct {
	val *T
}

// New creates a new Option.
func New[T any](val T) Option[T] {
	return Option[T]{val: &val}
}

// Some returns the value, if any, and a bool indicating if the value is non-nil.
func (o Option[T]) Some() (T, bool) {
	if o.val == nil {
		return *new(T), false
	}
	return *o.val, true
}

// None asserts if the value is nil.
func (o Option[T]) None() bool {
	return o.val == nil
}

// Unwrap extracts the value. It panics if the inner value is nil
func (o Option[T]) Unwrap() T {
	return *o.val
}

// UnwrapOr extracts the value. If the inner value is nil it returns the value specified by `val`.
func (o Option[T]) UnwrapOr(val T) T {
	if o.None() {
		return val
	}
	return *o.val
}

// UnwrapOrElse extracts the value. If the inner value is nil it
// executes the specified function and returns its value.
func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.None() {
		return f()
	}
	return *o.val
}

// UnwrapOrDefault extracts the value. It returns the type specific default value if the inner value is nil.
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
