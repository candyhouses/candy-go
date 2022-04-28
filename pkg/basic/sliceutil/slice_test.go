package sliceutil

import (
	"reflect"
	"testing"
)

func TestInster(t *testing.T) {

	type args[T int | string] struct {
		oldSlice []T
		i        int
		newSlice []T
	}
	type intTest struct {
		name string
		args args[int]
		want []int
	}
	int_tests := []intTest{

		{
			name: "add a slice",
			args: args[int]{
				oldSlice: []int{1, 2, 3, 4, 5},
				i:        1,
				newSlice: []int{5, 6, 7},
			},
			want: []int{1, 5, 6, 7, 2, 3, 4, 5},
		},
		{
			name: "add a element",
			args: args[int]{
				oldSlice: []int{1, 2, 3, 4, 5},
				i:        1,
				newSlice: []int{6},
			},
			want: []int{1, 6, 2, 3, 4, 5},
		},

		{
			name: "add to header element",
			args: args[int]{
				oldSlice: []int{1, 2, 3, 4, 5},
				i:        0,
				newSlice: []int{0},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
	}
	type strTest struct {
		name string
		args args[string]
		want []string
	}
	str_tests := []strTest{

		{
			name: "add a slice",
			args: args[string]{
				oldSlice: []string{"a", "b", "c", "d", "e"},
				i:        1,
				newSlice: []string{"i", "w"},
			},
			want: []string{"a", "i", "w", "b", "c", "d", "e"},
		},

		{
			name: "add a element",
			args: args[string]{
				oldSlice: []string{"a", "b", "c", "d", "e"},
				i:        0,
				newSlice: []string{"c"},
			},
			want: []string{"c", "a", "b", "c", "d", "e"},
		},
	}

	for _, tt := range int_tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.oldSlice, tt.args.i, tt.args.newSlice...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inster() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range str_tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.oldSlice, tt.args.i, tt.args.newSlice...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inster() = %v, want %v", got, tt.want)
			}
		})
	}
}
