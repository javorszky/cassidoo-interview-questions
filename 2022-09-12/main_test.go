package sept122022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cornerHit(t *testing.T) {
	type args struct {
		logoSize   dimension
		logoCoord  coordinate
		screenSize dimension
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "hits immediately",
			args: args{
				logoSize:   dimension{w: 4, h: 4},
				logoCoord:  coordinate{x: 0, y: 0},
				screenSize: dimension{w: 20, h: 20},
			},
			want:    true,
			want1:   1,
			wantErr: assert.NoError,
		},
		{
			name: "hits immediately from different starting position",
			args: args{
				logoSize:   dimension{w: 4, h: 4},
				logoCoord:  coordinate{x: 7, y: 7},
				screenSize: dimension{w: 20, h: 20},
			},
			want:    true,
			want1:   1,
			wantErr: assert.NoError,
		},
		{
			name: "hits after 1 bump",
			args: args{
				logoSize:   dimension{w: 4, h: 4},
				logoCoord:  coordinate{x: 0, y: 0},
				screenSize: dimension{w: 36, h: 20},
			},
			want:    true,
			want1:   2,
			wantErr: assert.NoError,
		},
		{
			name: "hits after 1 bump from different starting position",
			args: args{
				logoSize:   dimension{w: 4, h: 4},
				logoCoord:  coordinate{x: 8, y: 8},
				screenSize: dimension{w: 36, h: 20},
			},
			want:    true,
			want1:   2,
			wantErr: assert.NoError,
		},
		{
			name: "hits after 2 bumps",
			args: args{
				logoSize:   dimension{w: 4, h: 4},
				logoCoord:  coordinate{x: 0, y: 0},
				screenSize: dimension{w: 52, h: 20},
			},
			want:    true,
			want1:   3,
			wantErr: assert.NoError,
		},
		{
			name: "hits after 2 bumps from different starting position",
			args: args{
				logoSize:   dimension{w: 4, h: 4},
				logoCoord:  coordinate{x: 8, y: 8},
				screenSize: dimension{w: 52, h: 20},
			},
			want:    true,
			want1:   3,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := cornerHit(tt.args.logoSize, tt.args.logoCoord, tt.args.screenSize)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			tt.wantErr(t, err)
		})
	}
}
