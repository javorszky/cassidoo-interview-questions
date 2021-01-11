package oct52020

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortAndRotateSlice(t *testing.T) {
	type args struct {
		integers []int
		pivot    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sorts and pivots array",
			args: args{
				integers: []int{9, 5, 3, 2, 5},
				pivot:    2,
			},
			want: []int{5, 5, 9, 2, 3},
		},
		{
			name: "sorts and pivots array of 0 length",
			args: args{
				integers: []int{},
				pivot:    0,
			},
			want: []int{},
		},
		{
			name: "sorts and pivots slice around pivot 0",
			args: args{
				integers: []int{9, 5, 3, 2, 5},
				pivot:    0,
			},
			want: []int{2, 3, 5, 5, 9},
		},
		{
			name: "sorts and pivots slice around pivot max",
			args: args{
				integers: []int{9, 5, 3, 2, 5},
				pivot:    5,
			},
			want: []int{2, 3, 5, 5, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortAndRotateSlice(tt.args.integers, tt.args.pivot)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_lowestValue(t *testing.T) {
	type args struct {
		integers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "finds lowest in a boring slice",
			args: args{
				integers: []int{5, 9, 2, 3, 5},
			},
			want: 2,
		},
		{
			name: "finds lowest in a slice that was pivoted at element 0",
			args: args{
				integers: []int{2, 3, 5, 5, 9},
			},
			want: 2,
		},
		{
			name: "finds lowest in a slice where all elements are the same",
			args: args{
				integers: []int{2, 2, 2, 2, 2, 2, 2, 2, 2},
			},
			want: 2,
		},
		{
			name: "finds lowest in an empty slice",
			args: args{
				integers: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, lowestValue(tt.args.integers))
		})
	}
}

func Benchmark_lowestValue(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{
			name: "bench slice of 10",
			n:    10,
		},
		{
			name: "bench slice of 100",
			n:    100,
		},
		{
			name: "bench slice of 1,000",
			n:    1000,
		},
		{
			name: "bench slice of 10,000",
			n:    10000,
		},
		{
			name: "bench slice of 100,000",
			n:    100000,
		},
		{
			name: "bench slice of 1,000,000",
			n:    1000000,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			sl := generateSlices(b, bm.n)
			for i := 0; i < b.N; i++ {
				lowestValue(sl)
			}
		})
	}
}

func generateSlices(b *testing.B, n int) []int {
	b.Helper()
	slice := make([]int, 0, n)
	for i := 0; i < n; i++ {
		slice = append(slice, rand.Intn(n))
	}
	return sortAndRotateSlice(slice, 0)
}
