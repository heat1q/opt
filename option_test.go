package opt

import (
	"reflect"
	"testing"
)

func TestNewOption(t *testing.T) {
	s := "test"
	tests := []struct {
		name string
		args string
		want Option[string]
	}{
		{
			args: s,
			want: Option[string]{
				val: &s,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_Some_None(t *testing.T) {
	t.Run("value is present", func(t *testing.T) {
		o := New("potato")
		val, ok := o.Some()
		assertTrue(t, ok)
		assertFalse(t, o.None())
		assertEqual(t, "potato", val)
	})
	t.Run("value is not present", func(t *testing.T) {
		o := Option[string]{}
		val, ok := o.Some()
		assertFalse(t, ok)
		assertTrue(t, o.None())
		assertEqual(t, "", val)
	})
}

func TestOption_UnwrapOrDefault(t *testing.T) {
	t.Run("default string", func(t *testing.T) {
		var want string
		o := Option[string]{}
		assertEqual(t, want, o.UnwrapOrDefault())
	})
	t.Run("default int", func(t *testing.T) {
		var want int
		o := Option[int]{}
		assertEqual(t, want, o.UnwrapOrDefault())
	})
	t.Run("default bool", func(t *testing.T) {
		var want bool
		o := Option[bool]{}
		assertEqual(t, want, o.UnwrapOrDefault())
	})
	t.Run("default ptr", func(t *testing.T) {
		var want *int
		o := Option[*int]{}
		assertEqual(t, want, o.UnwrapOrDefault())
	})
	t.Run("default map", func(t *testing.T) {
		var want map[string]string
		o := Option[map[string]string]{}
		assertEqual(t, want, o.UnwrapOrDefault())
	})
	t.Run("default interface", func(t *testing.T) {
		var want any
		o := Option[any]{}
		assertEqual(t, want, o.UnwrapOrDefault())
	})
	t.Run("default struct", func(t *testing.T) {
		var want struct{}
		o := Option[struct{}]{}
		assertEqual(t, want, o.UnwrapOrDefault())
	})
}

func assertEqual(t *testing.T, expected, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal: expected = %v, actual = %v", expected, actual)
	}
}

func assertTrue(t *testing.T, expected bool) {
	if !expected {
		t.Errorf("Should be true")
	}
}

func assertFalse(t *testing.T, expected bool) {
	if expected {
		t.Errorf("Should be false")
	}
}
