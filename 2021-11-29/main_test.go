package november292021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_phoneLetter(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    []string
		wantErr bool
	}{
		{
			name:    "returns error for empty input string",
			s:       "",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "returns error when string contains a non-number",
			s:       "7434a",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "returns error when string contains a 0",
			s:       "8440",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "returns error when string contains a 1",
			s:       "8441",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := phoneLetter(tt.s)
			assert.ElementsMatch(t, got, tt.want)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
