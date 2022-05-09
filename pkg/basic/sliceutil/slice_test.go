package sliceutil

import (
	"reflect"
	"testing"
)

type Dog struct {
	kind string
	name string
}

func TestInster(t *testing.T) {

	type args[T any] struct {
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

	type stuTest struct {
		name string
		args args[Dog]
		want []Dog
	}

	stu_test := []stuTest{{
		name: "hello",
		args: args[Dog]{
			oldSlice: []Dog{Dog{name: "胖虎", kind: "柯基"}, Dog{name: "hi", kind: "tugou"}},
			i:        0,
			newSlice: []Dog{Dog{name: "add", kind: "test"}},
		},
		want: []Dog{Dog{name: "add", kind: "test"}, Dog{name: "胖虎", kind: "柯基"}, Dog{name: "hi", kind: "tugou"}},
	}}

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

	for _, tt := range stu_test {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.oldSlice, tt.args.i, tt.args.newSlice...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inster() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestDelete(t *testing.T) {
	type stringArgs struct {
		sli   []string
		start int
		end   int
	}
	tests := []struct {
		name string
		args stringArgs
		want []string
	}{

		{name: "normal case", args: stringArgs{sli: []string{"zero", "one", "two", "three", "fore"}, start: 1, end: 3}, want: []string{"zero", "fore"}},
		{name: "normal case", args: stringArgs{sli: []string{"zero", "one", "two", "three", "fore"}, start: 0, end: 4}, want: nil},
		{name: "normal case", args: stringArgs{sli: []string{"zero", "one", "two", "three", "fore"}, start: 0, end: 0}, want: []string{"one", "two", "three", "fore"}},
		{name: "normal case", args: stringArgs{sli: []string{"zero", "one", "two", "three", "fore"}, start: 4, end: 4}, want: []string{"zero", "one", "two", "three"}},
		{name: "normal case", args: stringArgs{sli: []string{"zero", "one", "two", "three", "fore"}, start: 4, end: 3}, want: []string{"zero", "one", "two", "three", "fore"}},
		{name: "normal case", args: stringArgs{sli: []string{"zero", "one", "two", "three", "fore"}, start: 1, end: 5}, want: []string{"zero", "one", "two", "three", "fore"}},
		{name: "normal case", args: stringArgs{sli: nil, start: 1, end: 3}, want: nil},
		{name: "normal case", args: stringArgs{sli: []string{}, start: 1, end: 3}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(tt.args.sli, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
