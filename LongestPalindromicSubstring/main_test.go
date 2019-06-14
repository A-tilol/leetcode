package longestpalindromicsubstring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "2",
			args: args{
				s: "abcbaabcba",
			},
			want: "abcbaabcba",
		},
		{
			name: "3",
			args: args{
				s: "abcbaabcbahoge",
			},
			want: "abcbaabcba",
		},
		{
			name: "4",
			args: args{
				s: "hogeabcbaabcba",
			},
			want: "abcbaabcba",
		},
		{
			name: "5",
			args: args{
				s: "hogeabcbaabcbahoge",
			},
			want: "abcbaabcba",
		},
		{
			name: "6",
			args: args{
				s: "hogeabccbahogeabcdcbahoge",
			},
			want: "abcdcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
