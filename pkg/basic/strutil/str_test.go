package strutil

import (
	"testing"
)

func TestIsBlank(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "not empty string", args: args{str: "asdd"}, want: false},
		{name: "not empty string start with blank char", args: args{str: "   asdd"}, want: false},
		{name: "empty string", args: args{str: ""}, want: true},
		{name: "blank string", args: args{str: "      "}, want: true},
		{name: "blank string by tabs", args: args{str: "\t"}, want: true},
		{name: "blank string by line break", args: args{str: "\n"}, want: true},
		{name: "blank string by line enter", args: args{str: "\r"}, want: true},
		{name: "blank string by all blank char ", args: args{str: "      \r \n  \t"}, want: true},
		{name: "chinese check ", args: args{str: " 你好"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.args.str); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasBlank(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "all not Blank", args: args{strs: []string{"sdfs", "asdfas", "asdfabads", "  ikdji", "\t asdifn", "\n asndfoaisdn", "\r asdhfoadbn"}}, want: false},
		{name: "one of array is Blank", args: args{strs: []string{"sdfs", " ", "asdfabads", "  ikdji", "\t asdifn", "\n asndfoaisdn", "\r asdhfoadbn"}}, want: true},
		{name: "all Blank", args: args{strs: []string{" ", " ", "\t"}}, want: true},
		{name: "fist is  Blank", args: args{strs: []string{" ", " asdf", "asdf\t"}}, want: true},
		{name: "last is  Blank", args: args{strs: []string{"asdf ", " asdf", "\t"}}, want: true},
		{name: "empty array", args: args{strs: []string{}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasBlank(tt.args.strs...); got != tt.want {
				t.Errorf("HasBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllBlank(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "all not Blank", args: args{strs: []string{"sdfs", "asdfas", "asdfabads", "  ikdji", "\t asdifn", "\n asndfoaisdn", "\r asdhfoadbn"}}, want: false},
		{name: "one of array is Blank", args: args{strs: []string{"sdfs", " ", "asdfabads", "  ikdji", "\t asdifn", "\n asndfoaisdn", "\r asdhfoadbn"}}, want: false},
		{name: "all Blank", args: args{strs: []string{" ", " ", "\t"}}, want: true},
		{name: "fist is  Blank", args: args{strs: []string{" ", " asdf", "asdf\t"}}, want: false},
		{name: "last is  Blank", args: args{strs: []string{"asdf ", " asdf", "\t"}}, want: false},
		{name: "empty array", args: args{strs: []string{}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllBlank(tt.args.strs...); got != tt.want {
				t.Errorf("IsAllBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlankToDefault(t *testing.T) {
	type args struct {
		str        string
		defaultStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "not empty string", args: args{str: "asdd", defaultStr: "fdhjrtu"}, want: "asdd"},
		{name: "empty string", args: args{str: "", defaultStr: "fdhjrtu"}, want: "fdhjrtu"},
		{name: "blank string", args: args{str: " ", defaultStr: "fdhjrtu"}, want: "fdhjrtu"},
		{name: "blank string", args: args{str: "\t", defaultStr: "fdhjrtu"}, want: "fdhjrtu"},
		{name: "blank string", args: args{str: "\n", defaultStr: "fdhjrtu"}, want: "fdhjrtu"},
		{name: "blank string", args: args{str: "\r", defaultStr: "fdhjrtu"}, want: "fdhjrtu"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlankToDefault(tt.args.str, tt.args.defaultStr); got != tt.want {
				t.Errorf("BlankToDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

//------- Empty

func TestIsEmpty(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "not empty string", args: args{str: "asdd"}, want: false},
		{name: "empty string", args: args{str: ""}, want: true},
		{name: "blank string", args: args{str: " "}, want: false},
		{name: "blank string", args: args{str: "\t"}, want: false},
		{name: "blank string", args: args{str: "\n"}, want: false},
		{name: "blank string", args: args{str: "\r"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.str); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmptyToDefault(t *testing.T) {
	type args struct {
		str        string
		defaultStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "not empty string", args: args{str: "asdd", defaultStr: "fdhjrtu"}, want: "asdd"},
		{name: "empty string", args: args{str: "", defaultStr: "fdhjrtu"}, want: "fdhjrtu"},
		{name: "blank string", args: args{str: " ", defaultStr: "fdhjrtu"}, want: " "},
		{name: "blank string", args: args{str: "\t", defaultStr: "fdhjrtu"}, want: "\t"},
		{name: "blank string", args: args{str: "\n", defaultStr: "fdhjrtu"}, want: "\n"},
		{name: "blank string", args: args{str: "\r", defaultStr: "fdhjrtu"}, want: "\r"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmptyToDefault(tt.args.str, tt.args.defaultStr); got != tt.want {
				t.Errorf("EmptyToDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasEmpty(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "all not Blank", args: args{strs: []string{"sdfs", "asdfas", "asdfabads", "  ikdji", "\t asdifn", "\n asndfoaisdn", "\r asdhfoadbn"}}, want: false},
		{name: "one of array is Blank", args: args{strs: []string{"sdfs", " ", "asdfabads", "  ikdji", "\t asdifn", "\n asndfoaisdn", "\r asdhfoadbn"}}, want: false},
		{name: "all Blank", args: args{strs: []string{" ", " ", "\t"}}, want: false},
		{name: "fist is  Blank", args: args{strs: []string{" ", " asdf", "asdf\t"}}, want: false},
		{name: "last is  Blank", args: args{strs: []string{"asdf ", " asdf", "\t"}}, want: false},
		{name: "one of array is Empty", args: args{strs: []string{"sdfs", "", "asdfabads", "  ikdji", "\t asdifn", "\n asndfoaisdn", "\r asdhfoadbn"}}, want: true},
		{name: "all Empty", args: args{strs: []string{"", "", ""}}, want: true},
		{name: "fist is  Empty", args: args{strs: []string{"", " asdf", "asdf\t"}}, want: true},
		{name: "last is  Empty", args: args{strs: []string{"asdf ", " asdf", ""}}, want: true},
		{name: "empty array", args: args{strs: []string{}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasEmpty(tt.args.strs...); got != tt.want {
				t.Errorf("HasEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllEmpty(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "all not Blank", args: args{strs: []string{"sdfs", "asdfas", "asdfabads", "  ikdji", "\t asdifn", "\n asndfoaisdn", "\r asdhfoadbn"}}, want: false},
		{name: "one of array is Blank", args: args{strs: []string{"sdfs", " ", "asdfabads", "  ikdji", "\t asdifn", "\n asndfoaisdn", "\r asdhfoadbn"}}, want: false},
		{name: "all Blank", args: args{strs: []string{" ", " ", "\t"}}, want: false},
		{name: "all Empty", args: args{strs: []string{"", "", ""}}, want: true},
		{name: "fist is  Blank", args: args{strs: []string{" ", " asdf", "asdf\t"}}, want: false},
		{name: "last is  Blank", args: args{strs: []string{"asdf ", " asdf", "\t"}}, want: false},
		{name: "fist is  Empty", args: args{strs: []string{"", " asdf", "asdf\t"}}, want: false},
		{name: "last is  Empty", args: args{strs: []string{"asdf ", " asdf", ""}}, want: false},
		{name: "empty array", args: args{strs: []string{}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllEmpty(tt.args.strs...); got != tt.want {
				t.Errorf("IsAllEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllNotEmpty(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "all not Blank", args: args{strs: []string{"sdfs", "asdfas", "asdfabads", "  ikdji", "\t asdifn", "\n asndfoaisdn", "\r asdhfoadbn"}}, want: true},
		{name: "all Empty", args: args{strs: []string{"", "", ""}}, want: false},
		{name: "one of strs is Empty", args: args{strs: []string{"", "a", "df"}}, want: false},
		{name: "one of strs is Blank", args: args{strs: []string{"\t", "a", "df"}}, want: true},
		{name: "empty array", args: args{strs: []string{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllNotEmpty(tt.args.strs...); got != tt.want {
				t.Errorf("IsAllNotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllNotBlank(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "all not Blank", args: args{strs: []string{"sdfs", "asdfas", "asdfabads", "  ikdji", "\t asdifn", "\n asndfoaisdn", "\r asdhfoadbn"}}, want: true},
		{name: "all Empty", args: args{strs: []string{"", "", ""}}, want: false},
		{name: "one of strs is Empty", args: args{strs: []string{"", "a", "df"}}, want: false},
		{name: "one of strs is Blank", args: args{strs: []string{"\t", "a", "df"}}, want: false},
		{name: "empty array", args: args{strs: []string{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllNotBlank(tt.args.strs...); got != tt.want {
				t.Errorf("IsAllNotBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNullOrUnderfined(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "null", args: args{str: "null"}, want: true},
		{name: "underfined", args: args{str: "undefined"}, want: true},
		{name: "error", args: args{str: "undefisdddddned"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNullOrUnderfined(tt.args.str); got != tt.want {
				t.Errorf("isNullOrUnderfined() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimBlankCharByMode(t *testing.T) {
	type args struct {
		str  string
		mode int
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{name: "left trim", args: args{str: "     hello", mode: -1}, want: "hello"},
		{name: "left trim but left don't have black char", args: args{str: "hello", mode: -1}, want: "hello"},
		{name: "left trim", args: args{str: "     hello   ", mode: -1}, want: "hello   "},
		{name: "left trim chinese", args: args{str: "     你好", mode: -1}, want: "你好"},
		{name: "right trim", args: args{str: "hello    ", mode: 1}, want: "hello"},
		{name: "right trim but right don't have black char", args: args{str: "hello", mode: 1}, want: "hello"},
		{name: "rigth trim", args: args{str: "     hello   ", mode: 1}, want: "     hello"},
		{name: "right trim chinese", args: args{str: "你好     ", mode: 1}, want: "你好"},
		{name: "all trim", args: args{str: "     hello", mode: 0}, want: "hello"},
		{name: "all trim but left don't have black char", args: args{str: "hello", mode: 0}, want: "hello"},
		{name: "all trim", args: args{str: "     hello   ", mode: 0}, want: "hello"},
		{name: "all trim chinese", args: args{str: "     你好", mode: 0}, want: "你好"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBlankCharByMode(tt.args.str, tt.args.mode); got != tt.want {
				t.Errorf("trimBlankCharByMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
