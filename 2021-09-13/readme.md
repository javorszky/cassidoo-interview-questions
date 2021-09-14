# 13th September

> This week’s question:
>
> Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
>
>Example:
> ```shell
> $ formParens(1)
> $ ["()"]
>
> $ formParens(3)
> $ ["((()))","(()())","(())()","()(())","()()()"]
> ```

## What I've done

Things I know to be true:
* Every paren combination starts with an opening paren and a closing one.
* I can't have more closing parens than opening parens at any point of the string from the left.
* There are equal number of opening and closing parens in a given combination.

With this, I created a recursive function that takes a partial paren combination, and then checks whether it can add an open, or a close paren, and return the result of them.

It can add an open paren if there are parens left to open. In the example, I can only open parens 3 times, so I call the recursive function, decrement the opens, and decrement the step, and increase the value.

It can close a paren if the value is above 0. Every open paren is a +1, every close paren is a -1. The recursive function keeps a running tally.

At the end, it returns the collection it did.

I also added a cute hill diagram to visualize the paren combinations. It looks like this with 5 pairs:

```shell
13th September 2021: all possible paren configs with 5 pairs of parens:
o----------o
|    ()    |
|   (  )   |
|  (    )  |
| (      ) |
|(        )|
o----------o
o----------o
|          |
|   ()()   |
|  (    )  |
| (      ) |
|(        )|
o----------o
o----------o
|          |
|   ()     |
|  (  )()  |
| (      ) |
|(        )|
o----------o
o----------o
|          |
|   ()     |
|  (  )    |
| (    )() |
|(        )|
o----------o
o----------o
|          |
|   ()     |
|  (  )    |
| (    )   |
|(      )()|
o----------o
o----------o
|          |
|     ()   |
|  ()(  )  |
| (      ) |
|(        )|
o----------o
o----------o
|          |
|          |
|  ()()()  |
| (      ) |
|(        )|
o----------o
o----------o
|          |
|          |
|  ()()    |
| (    )() |
|(        )|
o----------o
o----------o
|          |
|          |
|  ()()    |
| (    )   |
|(      )()|
o----------o
o----------o
|          |
|          |
|  ()  ()  |
| (  )(  ) |
|(        )|
o----------o
o----------o
|          |
|          |
|  ()      |
| (  )()() |
|(        )|
o----------o
o----------o
|          |
|          |
|  ()      |
| (  )()   |
|(      )()|
o----------o
o----------o
|          |
|          |
|  ()      |
| (  )  () |
|(    )(  )|
o----------o
o----------o
|          |
|          |
|  ()      |
| (  )     |
|(    )()()|
o----------o
o----------o
|          |
|     ()   |
|    (  )  |
| ()(    ) |
|(        )|
o----------o
o----------o
|          |
|          |
|    ()()  |
| ()(    ) |
|(        )|
o----------o
o----------o
|          |
|          |
|    ()    |
| ()(  )() |
|(        )|
o----------o
o----------o
|          |
|          |
|    ()    |
| ()(  )   |
|(      )()|
o----------o
o----------o
|          |
|          |
|      ()  |
| ()()(  ) |
|(        )|
o----------o
o----------o
|          |
|          |
|          |
| ()()()() |
|(        )|
o----------o
o----------o
|          |
|          |
|          |
| ()()()   |
|(      )()|
o----------o
o----------o
|          |
|          |
|          |
| ()()  () |
|(    )(  )|
o----------o
o----------o
|          |
|          |
|          |
| ()()     |
|(    )()()|
o----------o
o----------o
|          |
|          |
|      ()  |
| ()  (  ) |
|(  )(    )|
o----------o
o----------o
|          |
|          |
|          |
| ()  ()() |
|(  )(    )|
o----------o
o----------o
|          |
|          |
|          |
| ()  ()   |
|(  )(  )()|
o----------o
o----------o
|          |
|          |
|          |
| ()    () |
|(  )()(  )|
o----------o
o----------o
|          |
|          |
|          |
| ()       |
|(  )()()()|
o----------o
o----------o
|          |
|     ()   |
|    (  )  |
|   (    ) |
|()(      )|
o----------o
o----------o
|          |
|          |
|    ()()  |
|   (    ) |
|()(      )|
o----------o
o----------o
|          |
|          |
|    ()    |
|   (  )() |
|()(      )|
o----------o
o----------o
|          |
|          |
|    ()    |
|   (  )   |
|()(    )()|
o----------o
o----------o
|          |
|          |
|      ()  |
|   ()(  ) |
|()(      )|
o----------o
o----------o
|          |
|          |
|          |
|   ()()() |
|()(      )|
o----------o
o----------o
|          |
|          |
|          |
|   ()()   |
|()(    )()|
o----------o
o----------o
|          |
|          |
|          |
|   ()  () |
|()(  )(  )|
o----------o
o----------o
|          |
|          |
|          |
|   ()     |
|()(  )()()|
o----------o
o----------o
|          |
|          |
|      ()  |
|     (  ) |
|()()(    )|
o----------o
o----------o
|          |
|          |
|          |
|     ()() |
|()()(    )|
o----------o
o----------o
|          |
|          |
|          |
|     ()   |
|()()(  )()|
o----------o
o----------o
|          |
|          |
|          |
|       () |
|()()()(  )|
o----------o
o----------o
|          |
|          |
|          |
|          |
|()()()()()|
o----------o
```
