package october182021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reOrder(t *testing.T) {
	type args struct {
		arr     []string
		indices []int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "returns error when input slices are different lengths",
			args: args{
				arr:     []string{"a", "b"},
				indices: []int{1, 2, 3},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns error when indices contains duplicate value",
			args: args{
				arr:     []string{"a", "b"},
				indices: []int{1, 1},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns error when indices is out of bounds",
			args: args{
				arr:     []string{"a", "b", "c"},
				indices: []int{9, 10, 11},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "orders the strings correctly",
			args: args{
				arr:     []string{"C", "D", "E", "F", "G", "H"},
				indices: []int{3, 0, 4, 1, 2, 5},
			},
			want:    []string{"D", "F", "G", "C", "E", "H"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ordered, err := reOrder(tt.args.arr, tt.args.indices)
			assert.Equal(t, tt.want, ordered)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
