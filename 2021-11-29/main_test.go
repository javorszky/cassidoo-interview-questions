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

func Test_parseNumbers(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    []int
		wantErr bool
	}{
		{
			name:    "fails when input string has letters in it",
			s:       "1234a",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "correctly turns string of numbers into slice of ints",
			s:       "08472772",
			want:    []int{0, 8, 4, 7, 2, 7, 7, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseNumbers(tt.s)
			assert.Equal(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_shift(t *testing.T) {
	tests := []struct {
		name      string
		in        []int
		want      int
		want1     []int
		wantError bool
	}{
		{
			name:      "shifts one element off of the beginning of the array",
			in:        []int{1, 2, 3, 4, 5},
			want:      1,
			want1:     []int{2, 3, 4, 5},
			wantError: false,
		},
		{
			name:      "shift works fine with a slice of one length",
			in:        []int{4},
			want:      4,
			want1:     []int{},
			wantError: false,
		},
		{
			name:      "shift works fine with an initialized empty slice",
			in:        []int{},
			want:      0,
			want1:     nil,
			wantError: true,
		},
		{
			name:      "shift works fine with a nil slice",
			in:        nil,
			want:      0,
			want1:     nil,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := shift(tt.in)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			if tt.wantError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_assembler(t *testing.T) {
	type args struct {
		prefix  string
		numbers []int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "correctly returns slice for number 2",
			args: args{
				prefix:  "",
				numbers: []int{2},
			},
			want:    []string{"a", "b", "c"},
			wantErr: false,
		},
		{
			name: "correctly returns slice for number 9",
			args: args{
				prefix:  "",
				numbers: []int{9},
			},
			want:    []string{"w", "x", "y", "z"},
			wantErr: false,
		},
		{
			name: "correctly returns slice for number 4 with prefix",
			args: args{
				prefix:  "x",
				numbers: []int{4},
			},
			want:    []string{"xg", "xh", "xi"},
			wantErr: false,
		},
		{
			name: "correctly returns slice for numbers 3,5",
			args: args{
				prefix:  "",
				numbers: []int{3, 5},
			},
			want:    []string{"dj", "dk", "dl", "ej", "ek", "el", "fj", "fk", "fl"},
			wantErr: false,
		},
		{
			name: "correctly returns slice for numbers 6,7 with prefix",
			args: args{
				prefix:  "yo",
				numbers: []int{6, 7},
			},
			want: []string{"yomp", "yomq", "yomr", "yoms",
				"yonp", "yonq", "yonr", "yons",
				"yoop", "yooq", "yoor", "yoos"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := assembler(tt.args.prefix, tt.args.numbers)
			assert.ElementsMatch(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_phoneLetter1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "passes test with one number",
			args:    args{s: "9"},
			want:    []string{"w", "x", "y", "z"},
			wantErr: false,
		},
		{
			name:    "passes test with two numbers from the newsletter",
			args:    args{s: "23"},
			want:    []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
			wantErr: false,
		},
		{
			name:    "fails successfully with bad input (0 in the string)",
			args:    args{s: "220"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "fails successfully with bad input (1 in the string)",
			args:    args{s: "220"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "fails successfully with bad input (letter in the string)",
			args:    args{s: "22a"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := phoneLetter(tt.args.s)
			assert.ElementsMatch(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
