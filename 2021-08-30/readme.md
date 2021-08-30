# 30th August 2021

> This week’s question:
>
> Write a function to find the longest common prefix string in an array of strings.
>
>Example:
> ```shell
> $ longestPrefix(["cranberry","crawfish","crap"])
> $ "cra"
> 
> $ longestPrefix(["parrot", "poodle", "fish"])
> $ ""
> ```

## What I've done

Built a thing that builds a graph from the letters. Or a table. Basically it iterates over all the strings in the slice / array, and then splits them up to letters, and puts each letter into a map for the given position, so in the graph will look like this for the `cranberry, crawfish, crap` example:

```golang
g := graph{
	0: map[string]struct{}{
		"c": {},
	},
	1: map[string]struct{}{
		"r": {},
	},
	2: map[string]struct{}{
		"a": {},
	},
	3: map[string]struct{}{
		"n": {},
		"p": {},
		"w": {},
	},
	4: map[string]struct{}{
		"b": {},
		"f": {},
	},
	5: map[string]struct{}{
		"e": {},
		"i": {},
	},
	6: map[string]struct{}{
		"r": {},
		"s": {},
	},
	7: map[string]struct{}{
		"h": {},
		"r": {},
	},
	8: map[string]struct{}{
		"y": {},
	
	},
}
```
Position 0 has only the letter `c`, position 1 has only the letter `r`, position 2 has the letter `a`, position 3 has the letters `n, p, w`, etc...

Then I iterate over this graph, and if the number of letters in the position is 1, I add the letter to the prefix builder. The moment I encounter a position where the length is more than 1, or the graph ends, I return the collected string.

This approach should keep the complexity down to O(n) or so.
