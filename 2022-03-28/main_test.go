package march282022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containedItems(t *testing.T) {
	type args struct {
		in    string
		start int
		end   int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr func(*testing.T, error)
	}{
		{
			name: "negative start",
			args: args{
				in:    "doesnt matter",
				start: -9,
				end:   9,
			},
			want: 0,
			wantErr: func(t *testing.T, err error) {
				nse := negativeStartError
				assert.ErrorAs(t, err, &nse)
			},
		},
		{
			name: "out of bounds",
			args: args{
				in:    "123456",
				start: 0,
				end:   6,
			},
			want: 0,
			wantErr: func(t *testing.T, err error) {
				oobe := outOfBoundsError
				assert.ErrorAs(t, err, &oobe)
			},
		},
		{
			name: "bad string",
			args: args{
				in:    "0",
				start: 0,
				end:   1,
			},
			want: 0,
			wantErr: func(t *testing.T, err error) {
				bse := badStringError
				assert.ErrorAs(t, err, &bse)
			},
		},
		{
			name: "email example 1",
			args: args{
				in:    "|**|*|*",
				start: 0,
				end:   5,
			},
			want: 2,
			wantErr: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "email example 2",
			args: args{
				in:    "|**|*|*",
				start: 0,
				end:   6,
			},
			want: 3,
			wantErr: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "email example 3",
			args: args{
				in:    "|**|*|*",
				start: 1,
				end:   7,
			},
			want: 1,
			wantErr: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := containedItems(tt.args.in, tt.args.start, tt.args.end)

			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
