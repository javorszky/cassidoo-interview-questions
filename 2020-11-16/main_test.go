package nov162020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_specialPairs(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantNum int
	}{
		{
			name: "example in newsletter",
			args: args{
				arr: []int{1, 2, 3, 1, 1, 3},
			},
			want: [][]int{
				{0, 3},
				{0, 4},
				{3, 4},
				{2, 5},
			},
			wantNum: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairs := specialPairs(tt.args.arr)
			assert.ElementsMatch(t, tt.want, pairs)
			assert.Equal(t, tt.wantNum, numPairs(pairs))
		})
	}
}
