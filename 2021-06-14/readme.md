# 14th June 2021

> This week’s question:
> Given a direction, and a number of columns, write a function that outputs an arrow of asterisks (see the pattern in the examples below)!
>
> Example:
> ```bash
> $ printArrow('right', 3)
> Output:
> *
>  *
>   *
>  *
> *
>
> $ printArrow('left', 5)
> Output:
>     *
>    *
>   *
>  *
> *
>  *
>   *
>    *
>     *
>      *
> ```

## Approach

I opted for a recursive solution. The heart and soul of the solution is a setup function that takes the direction and the initial depth and sets up a _padder_ helper function which will give required amount of spaces at the beginning of the lines to the stars based on direction and current depth.

The recursive function takes this padder without needing to know what it does, and blindly uses it as it counts down _depth_ to zero.

The structure is that it prints a line of spaces and star, calls itself with depth-1, then prints another line of spaces and star at the same level.

If depth-1 is zero, it only prints one line of spaces and star without calling itself, so the stack can wind back up.
