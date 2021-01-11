# This week’s question:

Given an array of random integers, move all the zeros in the array to the end of the array. Try to keep this in O(n) time (or better)!

Example:

```
$ moveZeros([1, 2, 0, 1, 0, 0, 3, 6])
$ [1, 2, 1, 3, 6, 0, 0, 0]
```

From @cassidoo's [newsletter](https://buttondown.email/cassidoo/archive/if-you-spend-too-much-time-thinking-about-a-thing/)

## Benchmarks

```
go test -bench .
goos: darwin
goarch: amd64
BenchmarkMoveTenZeroes-16         	33359091	        32.8 ns/op
BenchmarkMoveHundredZeroes-16     	 5831109	       199 ns/op
BenchmarkMoveThousandZeroes-16    	  675056	      1798 ns/op
BenchmarkMove10kZeroes-16         	   73465	     16315 ns/op
BenchmarkMove100kZeroes-16        	    6838	    163073 ns/op
PASS
ok  	_/Users/<username>/Projects/cassidoo/17-aug-2020	6.244s
```