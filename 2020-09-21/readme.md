# code challenge 21st September 2019

This is from the newsletter here: https://buttondown.email/cassidoo/archive/real-change-enduring-change-happens-one-step-at-a/
 
##This week’s question:

Given an array of integers representing asteroids in a row, for each asteroid, the absolute value represents its size, and the sign represents its direction (positive = right, negative = left). Return the state of the asteroids after all collisions (assuming they are moving at the same speed). If two asteroids meet, the smaller one will explode. When they are the same size, they both explode. Asteroids moving in the same direction will never meet.

Example:
```bash
$ asteroids([5, 8, -5])
$ [5, 8] // The 8 and -5 collide, 8 wins. The 5 and 8 never collide.
```
```bash
$ asteroids([10, -10]) $ [] // The 10 and -10 collide and they both explode. 
```

## assumptions

The value 0 is always going to be a different sign to the previous asteroid. They always get destroyed. It doesn't matter whether the value used is `0` or `-0`.