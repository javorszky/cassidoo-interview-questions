# 28th March 2022

[Link to the online version of the email](https://buttondown.email/cassidoo/archive/frustration-is-fuel-that-can-lead-to-the/)

> Given a string that represents items as asterisks (*) and compartment walls as pipes (|), a start index, and an end index, return the number of items in a closed compartment.
>
> Example:
> ```shell
> let str = '|**|*|*'
>
> $ containedItems(str, 0, 5)
> 2
>
> $ containedItems(str, 0, 6)
> 3
>
> $ containedItems(str, 1, 7)
> 1
> ```

## Solution

Uh, it's a walker. I take in string, do a bunch of error checking like making sure the start index is not negative, the end index is not out of bounds, and then walk along the characters of the string starting at `start` and ending with `end`.

Counting begins the first time I encounter a `|` (hence the `began` flag ), and then I increment a counter for every `*`. On subsequent `|`s I drain the `*` counter into a holder, reset it to zero, and keep on iterating.

If I encounter a character that isn't a `*` or `|`, I return an error.

This solution is O(n). The two boundary checks are constant time, then at _worst_, the code would need to iterate over each element of the incoming string.

### Better solution

Step 1 is to walk through the entire string and extract where the `|` are in the string. Save them in a slice. That one is guaranteed to be sorted.

Step 2: for each set of start / end, use this slice of pipe indices as basis. The difference between two of them, minus 1, is how many enclosed items are there.

I think it's still O(n). The other way I can imagine is as I'm stepping through the string, I would need to loop through each index-pair which sounds a lot more work.
