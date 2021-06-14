package jun142021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testLeft5Arrow = `    *
   *
  *
 *
*
 *
  *
   *
    *`

const testRight4Arrow = `*
 *
  *
   *
  *
 *
*`

func Test_drawArrow(t *testing.T) {
	type args struct {
		direction string
		depth     int
	}
	tests := []struct {
		name        string
		args        args
		want        string
		shouldPanic bool
	}{
		{
			name: "panics when direction is not left or right",
			args: args{
				direction: "not right",
				depth:     90,
			},
			want:        "",
			shouldPanic: true,
		},
		{
			name: "draws left arrow correctly",
			args: args{
				direction: "left",
				depth:     5,
			},
			want:        testLeft5Arrow,
			shouldPanic: false,
		},
		{
			name: "draws right arrow correctly",
			args: args{
				direction: "right",
				depth:     4,
			},
			want:        testRight4Arrow,
			shouldPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				assert.Panics(t, func() { drawArrow(tt.args.direction, tt.args.depth) })
				return
			} else {
				assert.NotPanics(t, func() { drawArrow(tt.args.direction, tt.args.depth) })
			}

			assert.Equal(t, tt.want, drawArrow(tt.args.direction, tt.args.depth))
		})
	}
}
