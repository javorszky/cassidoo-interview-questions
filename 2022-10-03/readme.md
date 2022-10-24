# 3rd October 2022


[Link to the online version of the email](https://buttondown.email/cassidoo/archive/hope-is-a-gift-you-dont-have-to-surrender-a-power/)

> Given two integers, generate a “fibonacci-like” sequence of n digits (where the next number in the pattern is the sum of the previous two numbers). Extra credit: Given a sequence, determine if the sequence is “fibonacci-like”.
>
> Example:
> ```shell
> let n = 5
>
> > fibLike(10, 20, n)
> > [10, 20, 30, 50, 80]
>
> > fibLike(3, 7, n)
> > [3, 7, 10, 17, 27]
> ```

## Solution

The first solution is kind of boring, for loop that starts at 2, goes until n, does an `a, b = b, a+b`, does a few checks, and attaches the new b onto the end of the sequence.

The extra credit takes in a sequence, starts iterating on each element starting from the 3rd element, and checks whether that's the sum of the previous two, and if it isn't, returns false. Otherwise true.
