# 12th September 2022


[Link to the online version of the email](https://buttondown.email/cassidoo/archive/i-enjoy-the-time-passing-i-think-its-a-privilege/)

> Remember the [bouncing DVD logo](https://bouncingdvdlogo.com/)? Given the dimensions of the logo, its initial coordinates, and the size of a screen, write a function that will determine if its next collision will hit the corner of the screen. Assume it is initially moving southeast with a slope of -1. Extra credit, figure out how many bounces/collisions it will take to hit a corner!
>
> Example:
> ```js
> let dimensions = [5,5]
> let initialCoordinates = [0,0] // you decide which part of the logo the coords map to
> let screenSize = [100,100]
>
> cornerHit(dimensions, initialCoordinates, screenSize)
> true // in one collision
>
> cornerHit(dimensions, [45,70], [400,200])
> false
> ```

## Solution
