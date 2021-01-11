# cassidoo 5th October 2020

https://buttondown.email/cassidoo/archive/if-it-costs-you-your-peace-its-too-expensive/

> Given an array that was once sorted in ascending order is rotated at some pivot unknown to you beforehand (so [0,2,4,7,9] might become [7,9,0,2,4], for example). Find the minimum value in that array in O(n) or less.

## solution

See code, but tldr:

1. make sure the slice of integers is indeed sorted and rotated, see `sortAndRotateSlice`
2. assume that the lowest value of an empty slice is 0
3. store the first element as the "lowest" for all non-empty slices as a fallback value
4. loop through each value in the slice and the moment the current value is less than the previous value, return the current value
5. fall back to returning the first value if we didn't have a lower one

It runs in O(n) or less per benchmark, except for the 1,000,000 one which I guess has extra overhead due to GC or something.

## test and benchmark

```
go test -bench=. ./...
goos: linux
goarch: amd64
pkg: github.com/javorszky/cassidoo-5-oct-2020
Benchmark_lowestValue/bench_slice_of_10-4               149785039                7.90 ns/op
Benchmark_lowestValue/bench_slice_of_100-4              16555582                70.9 ns/op
Benchmark_lowestValue/bench_slice_of_1,000-4             1854050               639 ns/op
Benchmark_lowestValue/bench_slice_of_10,000-4             175762              6394 ns/op
Benchmark_lowestValue/bench_slice_of_100,000-4             16897             63979 ns/op
Benchmark_lowestValue/bench_slice_of_1,000,000-4            1202            914915 ns/op
PASS
ok      github.com/javorszky/cassidoo-5-oct-2020        11.140s
```