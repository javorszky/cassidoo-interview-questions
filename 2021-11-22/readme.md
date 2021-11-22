# 22nd November 2021

[Link to the online version of the email](https://buttondown.email/cassidoo/archive/almost-everything-will-work-again-if-you-unplug/)

> Given an array of strings, group the anagrams together in separate arrays. An anagram is a word or phrase formed by rearranging the letters of another word or phrase, using all the original letters exactly once.
>
> Example:
> ```shell
> $ groupAnagrams(["eat","tea","ten","poop","net","ate"])
> $ [["poop"],["net","ten"],["eat","tea","ate"]]
> ```

## Solution

This was fun! I took advantage of my math classes from way back when I learned about the properties of products of numbers from primary school. Or maybe high school... remember all those times you said "pfffft Mrs. Smith, when will we EVER use this in real life?!?!??!" Yeah, now. :D

Anyways, given that in a product it doesn't matter in which order you multiply numbers by, ie

```text
a × a × b × c
c × a × a × b
b × a × c × a
```
These will all produce the same exact result. This is part 1.

Part 2 is that any number can be broken up into their prime factors:
```text
7695864 | 2
3847932 | 2
1923966 | 2
 961983 | 3
 320661 | 3
 106887 | 3
  35629 | 11
   3239 | 41
     79 | 79
      1  
```
Because they're prime, a single factor cannot be changed out to any other factor.

What follows from these two is that any sequence of prime numbers in any order will always give the same end product.

Therefore I assigned the first 26 prime number to each letter code in the latin alphabet, then iterated over each word, calculated their prime product, and then grouped the words by that number using a map.

Then I just iterated over the grouping and collected all groups into a slice of a slice of strings.

### Considerations

I really wanted to make this solution work with slices only. Technically I still think I can, but given that the product of the primes can be rather large for long words, having a slice with hundreds of millions of empty values just so I can put one into an index and not get an out of bounds panic seems worse than using a map.
