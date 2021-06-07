package jun072021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lonelyNumber(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "all different numbers are calculated correctly",
			args: args{
				a: 2,
				b: 3,
				c: 4,
			},
			want: 24,
		},
		{
			name: "first two numbers same returns third number",
			args: args{
				a: 2,
				b: 2,
				c: 3,
			},
			want: 3,
		},
		{
			name: "last two numbers same, returns first number",
			args: args{
				a: 2,
				b: 3,
				c: 3,
			},
			want: 2,
		},
		{
			name: "first and last number same, returns middle number",
			args: args{
				a: 2,
				b: 3,
				c: 2,
			},
			want: 3,
		},
		{
			name: "all numbers same, returns 1",
			args: args{
				a: 3,
				b: 3,
				c: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, lonelyNumber(tt.args.a, tt.args.b, tt.args.c), tt.want)
		})
	}
}
