package aug312020

import (
	"fmt"
	"sort"
)

var sliceArtificial = []int{30, 35, 40, 45, 90, 20, 60}

type signal struct {
	buy    int
	sell   int
	bV     int
	sV     int
	profit int
}

// BuyDay will return the day to buy.
func (s *signal) BuyDay() int {
	return s.buy
}

// SellDay will return the day to sell.
func (s *signal) SellDay() int {
	return s.sell
}

// Profit will return the profit between the two.
func (s *signal) Profit() int {
	return s.profit
}

// String will return the human readable advice given the slice.
func (s *signal) String() string {
	return fmt.Sprintf("buy on day %d for %d, sell on day %d for %d for a profit of %d", s.buy+1, s.bV, s.sell+1, s.sV, s.profit)
}

func Task() {
	s := stockBuySell(slice100000)
	fmt.Printf("31st August 2020\n----------------\n%s\n\n", s.String())
}

func stockBuySell(prices []int) signal {
	var (
		candidates []signal
		changed    = false
		done       = false
		buy        = 0
		sell       = 0
		buyValue   = 99999
		sellValue  = 0
	)

	maxProfit := maxDiff(prices)

	for buyIndex, bV := range prices {
		if done {
			break
		}

		// The value considered for buy is higher than what we already have. No reason to check any further.
		if bV >= buyValue {
			continue
		}

		// if the value considered for buy is better than what we had, but it's past the sell date, store the previous
		// pair of values in a candidate, and reset the sell.
		if bV < buyValue && buyIndex > sell {
			candidates = append(candidates, signal{
				buy:    buy,
				sell:   sell,
				bV:     buyValue,
				sV:     sellValue,
				profit: sellValue - buyValue,
			})
			sell = buyIndex
			sellValue = 0
		}

		buyValue = bV
		buy = buyIndex

		for sellIndex, sV := range prices {
			changed = false
			// Iterate over the part of the slice that comes after the element we're looking at for "buy".
			if sellIndex <= buyIndex {
				continue
			}

			if sV > sellValue {
				sell = sellIndex
				sellValue = sV
			}

			// if we've reached peak profit, abort.
			if changed && sellValue-buyValue >= maxProfit {
				done = true
				break
			}
		}
	}

	candidates = append(candidates, signal{
		buy:    buy,
		bV:     buyValue,
		sell:   sell,
		sV:     sellValue,
		profit: sellValue - buyValue,
	})

	// Sort all available candidates by profit.
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i].profit > candidates[j].profit
	})

	// Return the one with the biggest profit.
	return candidates[0]
}

func maxDiff(in []int) int {
	// Remove duplicates
	normalized := unique(in)
	sort.Ints(normalized)
	return normalized[len(normalized)-1] - normalized[0]
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
