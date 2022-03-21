# 21st March 2022

[Link to the online version of the email](https://buttondown.email/cassidoo/archive/the-best-time-to-make-friends-is-before-you-need/)

> Given a list of times in a 24-hour period, return the smallest interval between two times in the list. Hint: you can do this in O(n) time!
>
> Example:
> ```shell
> $ smallestTimeInterval(['01:00', '08:15', '11:30', '13:45', '14:10', '20:05'])
> $ '25 minutes'
> ```

## Solution

1. Turn every string entry into an int of minutes. "04:33" becomes 273 (4 * 60 + 33)
2. Store every number in a slice
3. Order said slice
4. Walk through the ordered slice, checking the difference between current and previous, and compare against current smallest difference. If diff is smaller, replace smallest.
5. Return smallest difference
6. Convert back to string

Not O(n) yet because of the ordering, but that assumes that the input array is random, and not ordered. If the input array is ordered, then the sort can be left out, in which case it is an O(n) operation.
