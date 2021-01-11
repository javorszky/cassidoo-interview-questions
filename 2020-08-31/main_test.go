package aug312020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_stockBuySell10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stockBuySell(slice10)
	}
}

func Benchmark_stockBuySell100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stockBuySell(slice100)
	}
}

func Benchmark_stockBuySell1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stockBuySell(slice1000)
	}
}

func Benchmark_stockBuySell10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stockBuySell(slice10000)
	}
}

func Benchmark_stockBuySell100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stockBuySell(slice100000)
	}
}

func Test_stockBuySell(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   string
	}{
		{
			name:   "correctly figures out buy sell signal",
			prices: []int{110, 180, 260, 40, 310, 535, 695},
			want:   "buy on day 4 for 40, sell on day 7 for 695 for a profit of 655",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stockBuySell(tt.prices)
			assert.Equal(t, tt.want, s.String())
		})
	}
}
