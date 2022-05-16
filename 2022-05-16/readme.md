# 16th may 2022


[Link to the online version of the email](https://buttondown.email/cassidoo/archive/i-think-its-just-as-important-what-you-say-no-to/)

> Given two integer arrays of size n, return a new array of size n such that n consists of only unique elements and the sum of all its elements is maximum.
> 
> Example:
> ```js
> let arr1 = [7, 4, 10, 0, 1]
> let arr2 = [9, 7, 2, 3, 6]
> 
> $ maximizedArray(arr1, arr2)
> $ [9, 7, 6, 4, 10]
> ```

## Solution

Iterate over the two slices. Create a map that stores a value for each key. The numbers in the array are the keys. Have a fallback / buffer number read, set to 0 or minus however much is the minimum to guarantee any other number is going to be bigger.

For each index in the iteration over the array, grab both numbers, compare buffer number to both, and store the biggest one that isn't in the map yet.

Then grab the next biggest number of the remaining two that isn't in the map yet, and store that as the buffer number for the next iteration.

Repeat until finished with slice.

It will error out if it encounters the same number three times.
