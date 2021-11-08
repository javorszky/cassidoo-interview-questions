package november082021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_localPeaks(t *testing.T) {
	tests := []struct {
		name   string
		stream []int
		want   []int
	}{
		{
			name:   "finds local peak in small slice",
			stream: []int{1, 2, 1},
			want:   []int{1},
		},
		{
			name:   "finds a plateau",
			stream: []int{1, 2, 2, 1},
			want:   []int{1, 2},
		},
		{
			name:   "finds two peaks",
			stream: []int{1, 2, 1, 3, 1},
			want:   []int{1, 3},
		},
		{
			name:   "ignores end of slice plateau",
			stream: []int{1, 2, 1, 3, 3},
			want:   []int{1},
		},
		{
			name:   "finds two plateaus",
			stream: []int{1, 2, 2, 1, 3, 3, 2},
			want:   []int{1, 2, 4, 5},
		},
		{
			name:   "ignores beginning of slice peak",
			stream: []int{4, 3, 2, 3, 2},
			want:   []int{3},
		},
		{
			name:   "handles a plateau on ascent correctly",
			stream: []int{1, 2, 3, 3, 3, 4, 5, 4},
			want:   []int{6},
		},
		{
			name:   "newsletter test 1",
			stream: []int{1, 2, 3, 1},
			want:   []int{2},
		},
		{
			name:   "newsletter test 2",
			stream: []int{1, 3, 2, 3, 5, 6, 4},
			want:   []int{1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, localPeaks(tt.stream))
		})
	}
}
