# 16th August 2021

> This week’s question:
>
> Given a string s, return the longest palindromic substring in s.
>
>Example:
> ```shell
> $ pSubstring('babad')
> $ 'bab' // or 'aba'
> ```

## What I've done

It's a brute force approach with some optimizations.

1. break up the incoming string's letters into a slice
2. store the first letter in a variable that will house our longest found palindrome
3. starting at the first letter, try to find both an odd palindrome, and an even palindrome in a recursive function
4. the recursive function is: given a string, a down index, and an up index, find the letters wrapping that, check if it would stay a palindrome, glue it up, and then pass the new palindrome into the same function, and only return if the next step would stop being a palindrome

### Edge cases

1. palindrome is even, for example `abba`. In that case the center string is `""`, and the up and down index is `current - distance + 1` and `current + distance`
2. there are super palindromes, ie palindromes made up of palindromes, eg `abacdcaba`, where `aba` on both ends are also palindromes by themselves

### Optimization

If the already found longest palindrome is longer than the max possible length of any newly found palindrome, then don't bother looking.

For example in the string `mabbacdcabbaefghijkl`, once we found `abbacdcabba` (11), skip checking when the center of the palindrome would be `g` or later, because the longest possible palindrome centered on `g` could only be 11 long, which we already have. Longest possible centered on `h` is only 9, which we know is worse than what we have.
