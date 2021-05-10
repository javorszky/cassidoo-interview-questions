package may102021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sameDigits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "value of 1 will return true",
			n:    1,
			want: true,
		},
		{
			name: "value of 10 will return true",
			n:    1,
			want: true,
		},
		{
			name: "value of 3 will return false",
			n:    3,
			want: false,
		},
		{
			name: "value of 251894 will return true",
			n:    251894,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sameDigits(tt.n))
		})
	}
}

func Test_sameDigitsString(t *testing.T) {
	tests := []struct {
		name string
		n    string
		want bool
	}{
		{
			name: "very large number 1234567890 will return true",
			n:    "1234567890",
			want: true,
		},
		{
			name: "very large number 11223344556677665544332211 will return false",
			n:    "11223344556677665544332211",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sameDigitsString(tt.n))
		})
	}
}
