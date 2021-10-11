package october112021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isOdious(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "14 is odious",
			args: args{n: 14},
			want: true,
		},
		{
			name: "5 is not odious",
			args: args{n: 5},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isOdious(tt.args.n))
		})
	}
}
