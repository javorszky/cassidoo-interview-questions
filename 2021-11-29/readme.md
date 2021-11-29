# 22nd November 2021

[Link to the online version of the email](https://buttondown.email/cassidoo/archive/correction-does-much-but-encouragement-does-more/)

> Given a string containing digits from 2-9, return all possible letter combinations that the number could represent based on phone numbers/letters. For example, 2 could be a, b, or c, 3 could be d, e, or f, and so on.
>
> Example:
> ```shell
> $ phoneLetter('9')
> $ ['w', 'x', 'y', 'z']
> 
> $ phoneLetter('23')
> $ ['ad', 'ae', 'af', 'bd', 'be', 'bf', 'cd', 'ce', 'cf']
> ```

## Solution

Had a map of numbers -> letters. Did some checking to make sure we're only handling numbers 2-9 inclusive.

Then made a recursive function that shifts the current number off of the list, chooses one of the letters for the current number (and iterates for all of them), and runs the same process for the rest of the numbers, and then collects them in a slice.
