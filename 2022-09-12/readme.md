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

1. encode all properties into width / height and coordinate
2. check to see whether logo fits into the screen to begin with
3. designate the coordinate as the top left corner of the logo
4. create a virtual screen size so I can only deal with the top left corner. The size of this virtual screen is screen w - logo w and screen h - logo h
5. create an infinite grid of this virtual screen that way I don't need to deal with different directions, only bottom right. It will slice through the virtual screen edges
6. find the x position relative to the top left corner of the virtual screen by walking back until the y coordinate is 0. So if the starting position was [x: 4, y: 6], and the virtual screen is 34 wide, then the x position is going to be 32. (4 - 6 = -2, 34 - 2 = 32)
7. from there, plot where the diagonal line is going to intersect the next parallel top line. It's 32 + whatever the virtual screen height is, modulo virtual screen width
8. do this again and again, keeping track of the points. If we see 0, we have hit a corner! If we see a value we've seen already, we hit a loop that will never exit!
9. if we did hit a 0, count how many horizontal and vertical edges we've hit, minus how many we've hit finding the y=0 starting point
