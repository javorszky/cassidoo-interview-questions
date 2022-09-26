package sept262022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findOnes(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "finds 0 for 0",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "finds 1 for 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "finds 0 for 10",
			args: args{n: 10},
			want: 0,
		},
		{
			name: "finds 5 for 15",
			args: args{n: 15},
			want: 5,
		},
		{
			name: "finds 7 for 928748937294787",
			args: args{n: 928748937294787},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findOnes(tt.args.n))
		})
	}
}

func Test_findTens(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "finds 0 for 0",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "finds 0 for 1",
			args: args{n: 1},
			want: 0,
		},
		{
			name: "finds 1 for 10",
			args: args{n: 10},
			want: 1,
		},
		{
			name: "finds 1 for 15",
			args: args{n: 15},
			want: 1,
		},
		{
			name: "finds 5 for 956",
			args: args{n: 956},
			want: 5,
		},
		{
			name: "finds 8 for 928748937294787",
			args: args{n: 928748937294787},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTens(tt.args.n), "findTens(%v)", tt.args.n)
		})
	}
}

func Test_ordinal(t *testing.T) {
	tests := []struct {
		name string
		in   uint64
		want string
	}{
		{
			name: "zeroth",
			in:   0,
			want: "0th",
		},
		{
			name: "first",
			in:   1,
			want: "1st",
		},
		{
			name: "second",
			in:   2,
			want: "2nd",
		},
		{
			name: "third",
			in:   3,
			want: "3rd",
		},
		{
			name: "fourth",
			in:   4,
			want: "4th",
		},
		{
			name: "tenth",
			in:   10,
			want: "10th",
		},
		{
			name: "eleventh",
			in:   11,
			want: "11th",
		},
		{
			name: "twelfth",
			in:   12,
			want: "12th",
		},
		{
			name: "thirteenth",
			in:   13,
			want: "13th",
		},
		{
			name: "fourteenth",
			in:   14,
			want: "14th",
		},
		{
			name: "twentieth",
			in:   20,
			want: "20th",
		},
		{
			name: "twenty first",
			in:   21,
			want: "21st",
		},
		{
			name: "twenty second",
			in:   22,
			want: "22nd",
		},
		{
			name: "twenty third",
			in:   23,
			want: "23rd",
		},
		{
			name: "twenty fourth",
			in:   24,
			want: "24th",
		},
		{
			name: "one hundredth",
			in:   100,
			want: "100th",
		},
		{
			name: "one hundred first",
			in:   101,
			want: "101st",
		},
		{
			name: "one hundred second",
			in:   102,
			want: "102nd",
		},
		{
			name: "one hundred third",
			in:   103,
			want: "103rd",
		},
		{
			name: "one hundred fourth",
			in:   104,
			want: "104th",
		},
		{
			name: "one hundred tenth",
			in:   110,
			want: "110th",
		},
		{
			name: "one hundred eleventh",
			in:   111,
			want: "111th",
		},
		{
			name: "one hundred twelfth",
			in:   112,
			want: "112th",
		},
		{
			name: "one hundred thirteenth",
			in:   113,
			want: "113th",
		},
		{
			name: "one hundred fourteenth",
			in:   114,
			want: "114th",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ordinal(tt.in), "ordinal(%v)", tt.in)
		})
	}
}
