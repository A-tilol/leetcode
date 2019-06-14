package stringtointeger

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				str: "42",
			},
			want: 42,
		},
		{
			name: "2",
			args: args{
				str: "   -42",
			},
			want: -42,
		},
		{
			name: "3",
			args: args{
				str: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "4",
			args: args{
				str: "words and 987",
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remove_prefix_white_space(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "123",
			},
			want: "123",
		},
		{
			name: "2",
			args: args{
				s: "     12 3",
			},
			want: "12 3",
		},
		{
			name: "3",
			args: args{
				s: " ",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removePrefixWhiteSpace(tt.args.s); got != tt.want {
				t.Errorf("remove_prefix_white_space() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractValidSub(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "123 hoge",
			},
			want: "123",
		},
		{
			name: "2",
			args: args{
				s: "123hoge",
			},
			want: "123",
		},
		{
			name: "3",
			args: args{
				s: "10",
			},
			want: "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractValidSub(tt.args.s); got != tt.want {
				t.Errorf("extractValidSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
