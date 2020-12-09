# 7th december 2020

> This week’s question:
> Given two non-negative integers n1 and n2 represented as strings, return the product of n1 and n2, also represented
> as a string. Neither input will start with 0, and don’t just convert it to an integer and do the math that way.
>
> Examples:
> ```bash
> $ stringMultiply(“123”, “456”)
> $ “56088”
> ```

## what I've done

Created two tables: a sum table and a product table. They are maps, so `sum["7"]["8"] = "15"`, `multiplication["7"]["8"] = "56"`

And then modeled what I've learned in high school:
```bash
 123 * 456 = 56088
 ---
 492
  615
   738
 -----
 56088
```
Complete with carrys and everything. The best part is that it can multiple two arbitrarily large string numbers together.

I'm about 100% sure the code can be made better, but it works. And has tests.

## how to use it

You need to have go installed on your computer.

Then clone the repository, and run it with `go run main.go` from the root of the repository. That will run this week's solution.

At some point I'll add interactive cli to the runner too.
