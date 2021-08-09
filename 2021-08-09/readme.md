# 8th August 2021

> This week’s question:
>
> Given a grid size, and a set of mines (in pairs of rows and columns), generate the [Minesweeper](https://en.wikipedia.org/wiki/Minesweeper_(video_game)) grid for that set of mines. You can use asterisks for mines, and an x for 0!
>
>Example:
> ```shell
> let gridSize = 5
> let mines = [[1, 3], [3, 5], [2, 4]]
> 
> $ generateMinesweeper(gridSize, mines)
> 
> xxxxx
> 11xxx
> *21xx
> 2*21x
> 12*1x
> ```

## What I've done

Implemented custom types for the different things:
* `Coordinate` for the coordinates. It's an array with a strict 2 value length.
* `Field` which is the representation of the minesweeper field. It knows how big it is, and has a full list of its fields.
* `FieldType`, which is an iota enum. The values go from `mine`, to `one` ... `eight` denoting all possible values an individual coordinate can take up on the field.

All these to help me reduce unnecessary error checking, like "does the coordinate have too few values? Too many?", and the like.

This also means I can use the coordinate type itself as a map key in the map that has all the values of the field.

Because the field values are iotas, which are ints internally, I can increment them easily.

### Setup

The constructor function takes the grid size and the list of mines. It will first and foremost create an "empty" square field where all fields have the value of `zero`. This step prepopulates the internal map, which makes it super easy to do error checking when we try to add a mine that wouldn't be on the map because its coordinates fall outside.

This is an exponential step based on what the size is as it needs to generate x<sup>2</sup> values.

### Adding the mines

Then for each mine first I check whether it's on the field. Because we have an existing, complete, but empty field (if adding the first mine), this is fairly simple using `value, ok := f.fields[coordinate]`. If `ok` here is false, we know that the coordinate would not be on the field.

I do some error checking, for example adding a mine to a field with a mine already on it will result in an error, so duplicate coordinates are no bueno.

Then I calculate the coordinates for the adjacent fields for the mine: -1, 0, 1 in both directions, ignore itself, and grab what's on that field. The possibilities here are:

* adjacent field is not on the field: we're on edge somewhere, ignore, continue
* adjacent field is a mine: not going to increment the value, ignore, continue
* adjacent field is itself (0, 0), it's a mine, and we're not changing it, ignore, continue
* adjacent field is a number. Add one, store it back, and continue

### Stringify

This one was surprisingly straightforward. I summon a `strings.Builder`, loop through the coordinates of the field bound by `1` and `size`, grab the value of the field in those coordinates, and add the string representation of the fieldtype to the string builder.

After every row I add a newline, and before returning after the last row, I cut the last newline.
