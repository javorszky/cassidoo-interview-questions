package oct032022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fibLike(t *testing.T) {
	type args struct {
		a uint64
		b uint64
		n uint64
	}
	tests := []struct {
		name    string
		args    args
		want    []uint64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "simple",
			args: args{
				a: 1,
				b: 1,
				n: 3,
			},
			want:    []uint64{1, 1, 2},
			wantErr: assert.NoError,
		},
		{
			name: "too short sequence",
			args: args{
				a: 1,
				b: 1,
				n: 1,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "wraps around maybe",
			args: args{
				a: 5,
				b: 90,
				n: 300,
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fibLike(tt.args.a, tt.args.b, tt.args.n)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
