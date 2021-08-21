package august162021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pSubString(t *testing.T) {
	tests := []struct {
		name     string
		incoming string
		want     string
	}{
		{
			name:     "longest palindrome in 'abcdefghi' is the 'a'",
			incoming: "abcdefghi",
			want:     "a",
		},
		{
			name:     "longest palindrome of 'babad' is 'bab'",
			incoming: "babad",
			want:     "bab",
		},
		{
			name:     "longest palindrome of 'hellothererehtolleh' is itself",
			incoming: "hellothererehtolleh",
			want:     "hellothererehtolleh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, pSubString(tt.incoming))
		})
	}
}
