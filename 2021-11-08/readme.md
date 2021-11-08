# 8th November 2021

[Link to the online version of the email](https://buttondown.email/cassidoo/archive/we-cannot-direct-the-wind-but-we-can-adjust-the/)

> Given an array of integers, return the index of each local peak in the array. A “peak” element is an element that is greater than its neighbors.
>
> Example:
> ```shell
> $ localPeaks([1,2,3,1])
> $ [2]
> 
> $ localPeaks([1,3,2,3,5,6,4])
> $ [1, 5]
> ```

## Solution

**Note:** Strictly speaking an array like `[1, 2, 2, 1]` should not have a peak because none of the elements are greater than its neighbours, so I expanded the problem definition to also include plateaus where consecutive same elements surrounded by lower neighbours also count.

Super simple single pass iteration on the incoming array. We keep a track of the previous values, and if the current value is not smaller than the previous, their index to record where the peak / plateu is.

As we go, when we encounter a smaller value than previous, we drain the list of peak / plateau indexes into the permanent record of where they are, and at the end of it we return the permanent record.
