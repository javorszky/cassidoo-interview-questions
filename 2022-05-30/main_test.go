package may302022

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_allUnique(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string that is unique",
			args: args{str: "foxes"},
			want: true,
		},
		{
			name: "string that is not unique",
			args: args{
				str: "balloon",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, allUnique(tt.args.str),
				fmt.Sprintf("allUnique(%s), wanted %t, got %t", tt.args.str, tt.want, !tt.want))
		})
	}
}
