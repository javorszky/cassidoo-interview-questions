# 7th June 2021

> This week’s question:
> Given three numbers, return their product. But, if one of the numbers is the same as another, it does not count: If two numbers are similar, return the lonely number. If all numbers are same, return 1.
>
> Example:
> ```bash
> $ lonelyNumber(1,2,3)
> $ 6
> 
> $ lonelyNumber(6,6,4)
> $ 4
> 
> $ lonelyNumber(7,7,7)
> $ 1
> ```

## Approach

Collect the three numbers into a map with the numbers being key, and the occurence being the value. For each number, increment the value of the map with the key of the number.

Then I start with a number of 1 as a basis, then iterate over the map, and for each key (the numbers passed in), where the value is larger than 1 (ie there's more than one of them), I skip counting. Otherwise I will multiply the previous product with the current number (so the first iteration is going to be 1 (starter product) * whatever number).

At the end I return the product.
