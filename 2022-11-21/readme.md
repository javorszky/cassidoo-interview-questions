# 21st November 2022

[Link to the online version of the email](https://buttondown.email/cassidoo/archive/normal-is-not-something-to-aspire-to-its-4437/)

> Write a function that takes a string of slashes (\ and /) and returns all of those slashes drawn downwards in a line connecting them.
>
> Example:
> ```shell
> $ verticalSlashes('\\\//\/\\')
> \
>  \
>   \
>   /
>  /
>  \
>  /
>  \
>   \
> $ verticalSlashes('\\\\')
> \
>  \
>   \
>    \
> ```

## Solution

It's a for loop. Each character goes into a new line, so the only thing I have to figure out is how much padding to apply so the slash / backslash are the correct number of characters to the right.

Each \ gets a +1, each / gets a -1, I start at 0 and end with an array / slice of numbers.

Then I normalize that, which is essentially shifting every value by the lowest number to the right, so all the numbers are zero or more.

Then I output the characters with that much padding, and some minor adjustment for the / lines where I have to add one more to the spaces.
