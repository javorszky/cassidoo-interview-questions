# 5th September 2022


[Link to the online version of the email](https://buttondown.email/cassidoo/archive/if-everything-was-perfect-you-would-never-learn/)

> Write a function `fromTo` that produces a generator, that will produce values in a range.
>
> Example:
> ```js
> let gen = fromTo(5,7)
> > gen()
> 5
> > gen()
> 6
> > gen()
> 7
> > gen()
> undefined
> ```

## Solution

It's currying. Create a function that sets up one variable that will live for as long as the generator function live that tracks how many times the function was called.

Then create the generator function, check whether the start + number of called is bigger than the end. If so, return error, if not, return the next number, and then increment the called counter.
