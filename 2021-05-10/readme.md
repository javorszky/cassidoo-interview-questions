# 10th May 2021

> This week’s question:
>
> Given an integer n, return true if n^3 and n have the same set of digits.
>
> Test cases:
> ```bash
> $ sameDigits(1) // true
> $ sameDigits(10) // true
> $ sameDigits(251894) // true
> $ sameDigits(251895) // false
> ```

## what I've done

Reused the solution from [december 7, 2020](../2020-12-07/readme.md) which had has multiply two numbers as strings, which circumvents the whole "integer overflow" issue.

And then turn the numbers into strings, if they aren't strings to begin with, then create a map of the digits, and compare the cubed number, and bail the moment it has a number that isn't in a map.
