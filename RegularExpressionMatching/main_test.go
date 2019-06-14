package regularexpressionmatching

import (
	"reflect"
	"testing"
)

func Test_getNomalSubs(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				s: "hoge",
			},
			want: []string{
				"hoge",
			},
		},
		{
			name: "2",
			args: args{
				s: "hoge.piyo",
			},
			want: []string{
				"hoge",
				"piyo",
			},
		},
		{
			name: "3",
			args: args{
				s: "hoge.*piyo",
			},
			want: []string{
				"hoge",
				"piyo",
			},
		},
		{
			name: "4",
			args: args{
				s: ".*hoge.*piyo**",
			},
			want: []string{
				"hoge",
				"piyo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNomalSubs(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNomalSubs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
