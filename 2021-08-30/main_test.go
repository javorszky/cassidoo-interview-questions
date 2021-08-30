package august302021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPrefix(t *testing.T) {
	type args struct {
		arrayOfStrings []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "finds empty string for array of strings where there's no common prefix",
			args: args{
				arrayOfStrings: []string{
					"parrot",
					"poodle",
					"fish",
				},
			},
			want: "",
		},
		{
			name: "finds longest string if one of the words is contained by the others",
			args: args{
				arrayOfStrings: []string{
					"pre",
					"premonition",
					"predate",
					"predator",
				},
			},
			want: "pre",
		},
		{
			name: "finds common prefix where all of the words are longer",
			args: args{
				arrayOfStrings: []string{
					"cranberry",
					"crawfish",
					"crap",
				},
			},
			want: "cra",
		},
		{
			name: "collects the entire word if the input is just one word",
			args: args{
				arrayOfStrings: []string{
					"solitary",
				},
			},
			want: "solitary",
		},
		{
			name: "collects the entire word if the input is just one word multiple times",
			args: args{
				arrayOfStrings: []string{
					"duplicates",
					"duplicates",
					"duplicates",
				},
			},
			want: "duplicates",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestPrefix(tt.args.arrayOfStrings))
		})
	}
}
