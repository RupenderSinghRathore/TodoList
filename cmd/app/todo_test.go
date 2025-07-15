package main

import (
	"reflect"
	"testing"
)

func TestCleanUp(t *testing.T) {
	tests := []struct {
		name  string
		input *Todos
		want  *Todos
	}{
		{
			input: &Todos{&Todo{}, nil, &Todo{}, nil, &Todo{}, nil, nil},
			want:  &Todos{&Todo{}, &Todo{}, &Todo{}},
		},
	}

	for _, tt := range tests {
		cleanup(tt.input)
		if !reflect.DeepEqual(tt.input, tt.want) {
			t.Errorf("got: %v   want: %v\n", tt.input, tt.want)
		}
	}
}
