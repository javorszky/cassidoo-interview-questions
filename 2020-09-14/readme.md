# Code challenge 14th September 2020

Given an array of increasing integers, find the length of the longest fibonacci-like subsequence of the array. If one does not exist, return 0. A sequence is “fibonacci-like” if X_i + X_{i+1} = X_{i+2}.

Example:
```
$ fibonacciLike([1,3,7,11,12,14,18])
$ 3 // these sequences: [1,11,12], [3,11,14] or [7,11,18]
```