package dec282020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_find2020(t *testing.T) {
	type args struct {
		s string
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, find2020(tt.args.s))
		})
	}
}
