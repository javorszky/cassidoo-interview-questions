package dec72020

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringMultiply(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "multiplies two single digit numbers together without overflow",
			args: args{
				a: "3",
				b: "2",
			},
			want: "6",
		},
		{
			name: "multiplies two single digit numbers together without overflow the other way around",
			args: args{
				a: "2",
				b: "3",
			},
			want: "6",
		},
		{
			name: "multiplies two single digit numbers together with overflow",
			args: args{
				a: "3",
				b: "9",
			},
			want: "27",
		},
		{
			name: "multiplies two single digit numbers together with overflow the other way around",
			args: args{
				a: "9",
				b: "3",
			},
			want: "27",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, StringMultiply(tt.args.a, tt.args.b))
		})
	}
}

func Test_reverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverses string",
			args: args{
				s: "hello",
			},
			want: "olleh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseString(tt.args.s))
		})
	}
}

func Test_rotateMatrix(t *testing.T) {
	type args struct {
		input [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "can rotate matrix",
			args: args{
				input: [][]string{
					{"8", "3", "7", "", "", ""},
					{"", "5", "1", "6", "", ""},
					{"", "", "2", "9", "4", ""},
				},
			},
			want: [][]string{
				{"8", "", ""},
				{"3", "5", ""},
				{"7", "1", "2"},
				{"", "6", "9"},
				{"", "", "4"},
				{"", "", ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateMatrix(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addNumbers(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "adds two random digit numbers",
			args: args{
				a: []string{"34222", "99972661"},
			},
			want: "100006883",
		},
		{
			name: "adds two two digit numbers",
			args: args{
				a: []string{"34", "99"},
			},
			want: "133",
		},
		{
			name: "adds lots of one digit numbers where the result is a three digit number",
			args: args{
				a: []string{"9", "9", "9", "9", "9", "9", "9", "9", "9", "9", "9", "9"},
			},
			want: "108",
		},
		{
			name: "treats empty string as 0 value",
			args: args{
				a: []string{"9", "9", "", ""},
			},
			want: "18",
		},
		{
			name: "treats empty string as 0 value",
			args: args{
				a: []string{"9", "", ""},
			},
			want: "9",
		},
		{
			name: "treats all empty strings as 0 value",
			args: args{
				a: []string{"", "", ""},
			},
			want: "0",
		},
		{
			name: "treats all empty strings as 0 value when there's a 0 value present as well",
			args: args{
				a: []string{"", "", "0"},
			},
			want: "0",
		},

		//[]string{"0", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, addNumbers(tt.args.a...))
		})
	}
}

func Test_convertEmptyStringsToZeroes(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "replaces empty string with string 0 in slice",
			args: args{
				nums: []string{"84", "93", "", "", "", "4"},
			},
			want: []string{"84", "93", "0", "0", "0", "4"},
		},
		{
			name: "replaces empty string with string 0 in slice where only the last is a non empty string",
			args: args{
				nums: []string{"", "", "", "", "", "4"},
			},
			want: []string{"0", "0", "0", "0", "0", "4"},
		},
		{
			name: "replaces empty string with string 0 in slice with only empty strings",
			args: args{
				nums: []string{"", "", "", "", "", ""},
			},
			want: []string{"0", "0", "0", "0", "0", "0"},
		},
		{
			name: "replaces empty string with string 0 in slice where only the first is non empty string",
			args: args{
				nums: []string{"4", "", "", "", "", ""},
			},
			want: []string{"4", "0", "0", "0", "0", "0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, convertEmptyStringsToZeroes(tt.args.nums))
		})
	}
}

func Test_carryInnerOverflow(t *testing.T) {
	type args struct {
		iof []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "successfully carries over a long innerOverflow",
			args: args{
				iof: []string{"12"},
			},
			want: []string{"2", "1"},
		},
		{
			name: "successfully carries over a long innerOverflow",
			args: args{
				iof: []string{"12", "9", "9", "9", "9", "9"},
			},
			want: []string{"2", "0", "0", "0", "0", "0", "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, carryInnerOverflow(tt.args.iof))
		})
	}
}
