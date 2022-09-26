# 26th September 2022


[Link to the online version of the email]https://buttondown.email/cassidoo/archive/every-day-brings-new-choices-martha-beck/)

> Write a function to output the [ordinal](https://en.wikipedia.org/wiki/Ordinal_numeral) suffix of a positive integer concatenated to an inputted number.
>
> Example:
> ```shell
> > ordinal(3)
> > '3rd'
>
> > ordinal(57)
> > '57th'
> ```

## Solution

Take number, get the value for the ones and the tens. Use modulo operator to figure out the value. Because the inputs are `uint64` to enforce positive integerness and large enough space and make it easy to deal with `FormatUint` method, I can divide that by 10 which will discard the ones, so I can modulo that number to get the tens.

In English, rules for ordinals are:
* every number that ends in a 1 gets a "st" suffix
* except numbers that end in 11, which get a "th" suffix
* numbers that end in 2 get a "nd" suffix
* except numbers that end in 12, they get "th"
* numbers that end in 3 get a "rd" suffix
* except numbers that end in 13, they get "th"

Switch rules:
* if the value of tens is 1, get a th
* if the value of ones is 1, get a st
* if the value of ones is 2, get a nd
* if the value of ones is 3, get a rd
* everything else gets a th
