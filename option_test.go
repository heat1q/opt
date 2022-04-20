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
