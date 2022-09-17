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
		{
			name: "does not hit, loops",
			args: args{
				logoSize:   dimension{w: 4, h: 4},
				logoCoord:  coordinate{x: 1, y: 3},
				screenSize: dimension{w: 30, h: 30},
			},
			want:    false,
			want1:   0,
			wantErr: assert.NoError,
		},
		{
			name: "errors out, height is too big",
			args: args{
				logoSize:   dimension{w: 10, h: 10},
				logoCoord:  coordinate{x: 0, y: 0},
				screenSize: dimension{w: 14, h: 8},
			},
			want:    false,
			want1:   0,
			wantErr: assert.Error,
		},
		{
			name: "errors out, width is too big",
			args: args{
				logoSize:   dimension{w: 15, h: 2},
				logoCoord:  coordinate{x: 0, y: 0},
				screenSize: dimension{w: 14, h: 8},
			},
			want:    false,
			want1:   0,
			wantErr: assert.Error,
		},
		{
			name: "errors out, both values are too big",
			args: args{
				logoSize:   dimension{w: 10, h: 10},
				logoCoord:  coordinate{x: 0, y: 0},
				screenSize: dimension{w: 8, h: 8},
			},
			want:    false,
			want1:   0,
			wantErr: assert.Error,
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
