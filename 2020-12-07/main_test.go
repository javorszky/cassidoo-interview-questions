package dec72020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringMultiply(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "multiplies two single digit numbers together without overflow",
			args: args{
				a: "3",
				b: "2",
			},
			want: "6",
		},
		{
			name: "multiplies two single digit numbers together without overflow the other way around",
			args: args{
				a: "2",
				b: "3",
			},
			want: "6",
		},
		{
			name: "multiplies two single digit numbers together with overflow",
			args: args{
				a: "3",
				b: "9",
			},
			want: "27",
		},
		{
			name: "multiplies two single digit numbers together with overflow the other way around",
			args: args{
				a: "9",
				b: "3",
			},
			want: "27",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, stringMultiply(tt.args.a, tt.args.b))
		})
	}
}
