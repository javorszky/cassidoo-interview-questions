# 18th October 2021

[Link to the online version of the email](https://buttondown.email/cassidoo/archive/reality-is-one-of-the-possibilities-i-cannot/)

> Given an array of objects A, and an array of indexes B, reorder the objects in array A with the given indexes in array B.
> 
> Example:
> ```shell
> let a = [C, D, E, F, G, H];
> let b = [3, 0, 4, 1, 2, 5];
> 
> $ reorder(a, b) // a is now [D, F, G, C, E, H]
> ```

## Solution

The bulk of the code is error checking to be honest.

1. I need to make sure that the integers and the letters are the same length
2. then I shove the integer/letter pairings into a map that maps the value of the integer to the letter
3. then I check whether the length of that map is the same as the length of the array of integers. If it isn't, there was a duplicate number in the arrays
4. then I create a placeholder slice that should hold the letter in the correct order, per the input array, that has the same length and capacity as the array of integers
5. then I iterate over the map, and for each key of the map (the index after sorting), I put the value of the map at that point (the letter) into the placeholder array
6. if there's an out of bounds issue, the code would panic (meaning we skipped a number between 0 and n, or started at 1, or whatever), but the deferred recover call will catch that, set the named return and error, and release the code back to the wild
