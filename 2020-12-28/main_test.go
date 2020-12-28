package dec282020

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testString = "2220000202220020200"

var result int

func Test_find2020(t *testing.T) {
	type args struct {
		s string
	}
	funcs := []struct {
		name string
		f    func(string) int
	}{
		{
			name: "iterate over each string char, check first, then check rest",
			f:    find2020,
		},
		{
			name: "iterate over each string char, then check full slice of current:current+3",
			f:    find2020full,
		},
		{
			name: "iterate over byte range, check first, then check +1, +2, +3",
			f:    find2020byterange,
		},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns -1 for a string that's too short",
			args: args{
				s: "202",
			},
			want: -1,
		},
		{
			name: "returns -1 for a string that does not have 2020 in it",
			args: args{
				s: "00000200002222220000000022200002",
			},
			want: -1,
		},
		{
			name: "returns 1 for string that has multiple 2020s in it, the first one at 1",
			args: args{
				s: "0202000022222020000000",
			},
			want: 1,
		},
		{
			name: "returns 3 for string that has multiple 2020s in it, the first one at 3",
			args: args{
				s: "200202000022222020000000",
			},
			want: 3,
		},
		{
			name: "returns 16 for string that has one 2020 in it at the very end",
			args: args{
				s: "00000000000000002020",
			},
			want: 16,
		},
	}
	for _, tf := range funcs {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s: %s", tf.name, tt.name), func(t *testing.T) {
				assert.Equal(t, tt.want, tf.f(tt.args.s))
			})
		}
	}
}

func BenchmarkFind2020(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = find2020(testString)
	}
	result = r
}

func BenchmarkFind2020Full(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = find2020full(testString)
	}
	result = r
}

func BenchmarkFind2020byteslice(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = find2020byterange(testString)
	}
	result = r
}
