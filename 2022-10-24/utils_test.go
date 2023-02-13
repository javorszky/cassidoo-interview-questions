package oct242022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitSlice(t *testing.T) {
	tests := []struct {
		name  string
		s     []uint
		lean  side
		want  []uint
		want1 uint
		want2 []uint
	}{
		{
			name:  "splits a 5 long correctly left side",
			s:     []uint{1, 2, 3, 4, 5},
			lean:  leftSide,
			want:  []uint{1, 2},
			want1: 3,
			want2: []uint{4, 5},
		},
		{
			name:  "splits a 5 long correctly right side",
			s:     []uint{1, 2, 3, 4, 5},
			lean:  rightSide,
			want:  []uint{1, 2},
			want1: 3,
			want2: []uint{4, 5},
		},
		{
			name:  "splits a 4 long correctly, left leaning",
			s:     []uint{1, 2, 3, 4},
			lean:  leftSide,
			want:  []uint{1, 2},
			want1: 3,
			want2: []uint{4},
		},
		{
			name:  "splits a 4 long correctly, right leaning",
			s:     []uint{1, 2, 3, 4},
			lean:  rightSide,
			want:  []uint{1},
			want1: 2,
			want2: []uint{3, 4},
		},
		{
			name:  "splits a 2 long correctly left leaning",
			s:     []uint{1, 2},
			lean:  leftSide,
			want:  []uint{1},
			want1: 2,
			want2: []uint{},
		},
		{
			name:  "splits a 2 long correctly right leaning",
			s:     []uint{1, 2},
			lean:  rightSide,
			want:  []uint{},
			want1: 1,
			want2: []uint{2},
		},
		{
			name:  "splits a 1 long correctly left",
			s:     []uint{1},
			lean:  leftSide,
			want:  []uint{},
			want1: 1,
			want2: []uint{},
		},
		{
			name:  "splits a 1 long correctly right",
			s:     []uint{1},
			lean:  rightSide,
			want:  []uint{},
			want1: 1,
			want2: []uint{},
		},
		{
			name:  "splits a 0 long correctly",
			s:     []uint{},
			want:  []uint{},
			want1: 0,
			want2: []uint{},
		},
		{
			name:  "splits a nil correctly",
			s:     nil,
			want:  []uint{},
			want1: 0,
			want2: []uint{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := splitSlice(tt.s, tt.lean)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, got2)
		})
	}
}
