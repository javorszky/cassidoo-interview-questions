#28th December 2020

> This week’s question:
>
> You’re given a string of characters that are only 2s and 0s. Return the index of the first occurrence of “2020” without using the indexOf (or similar) function, and -1 if it’s not found in the string.
>
> Example:
> ```bash
> $ find2020(‘2220000202220020200’)
> $ 14
> ```

## what I've done

This is remarkably similar to [advent of code 2020 day 20 task 2](https://adventofcode.com/2020/day/20) where we had an image and had to find a "monster" in the image.

1. I check if the string is shorter than 4 characters. That automatically gets -1.
1. I loop through each character of the input string starting with the very first one, ending with the 4th from the end. I'm then looking at part of the string from `index` to `index+3`.
    1. If that spells out "2020", then I return whatever index I'm currently on.
    2. Otherwise I continue on with the loop.
1. If I haven't returned when I reach the end of the string and break out the loop, I return -1.

I don't check the string that it's actually only 2s and 0s. All I'm interested in is the "2" immediately followed by "020", so technically a string like "Gabor for president 2020" should also match it.

I also checked for three different implementations, and benchmarked them. The fastest is the `find2020full` with 9.85ns/op.
/
See this email online here: https://buttondown.email/cassidoo/archive/we-are-what-we-repeatedly-do-excellence-then-is/

