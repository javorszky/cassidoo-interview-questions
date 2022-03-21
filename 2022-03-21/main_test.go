package march212022

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numberToTimeString(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "formats only minutes",
			args: args{in: 45},
			want: "45 minutes",
		},
		{
			name: "formats hours and minutes",
			args: args{
				in: 61,
			},
			want: "1 hours and 1 minutes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, numberToTimeString(tt.args.in), tt.want)
		})
	}
}

func Test_entryToNumber(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "converts entry to number",
			args:    args{in: "01:12"},
			want:    72,
			wantErr: assert.NoError,
		},
		{
			name:    "less than an hour",
			args:    args{in: "00:03"},
			want:    3,
			wantErr: assert.NoError,
		},
		{
			name:    "throws error on bad",
			args:    args{in: "not a number"},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := entryToNumber(tt.args.in)
			if !tt.wantErr(t, err, fmt.Sprintf("entryToNumber(%v)", tt.args.in)) {
				return
			}
			assert.Equalf(t, tt.want, got, "entryToNumber(%v)", tt.args.in)
		})
	}
}
