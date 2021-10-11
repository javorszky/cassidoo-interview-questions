# 11th October 2021

> An “odious number” is a non-negative number that has an odd number of 1s in its binary expansion. Write a function that returns true if a given number is odious.
>
> Example:
> ```shell
> $ isOdious(14)
> $ true
> 
> $ isOdious(5)
> $ false
> ```

## Solution

Count (using core library `strings.Count` method) the number of `1`s in the binary number representation (got by using core library `fmt.Sprintf` with `%b` placeholder against an int) of the incoming number, and see whether after dividing by 2, the remainder is 1, or 0.
