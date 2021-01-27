## Part 1 notes

- The docking package holds data structures for bit manipulation.
- The Parser does what its name hopefully suggests. I elected to use strings as the memory's keys, because parsing uint64 didn't seem worth it. After all, the keys are only that: keys to a map. No computation is needed on them.
- Corner/invalid cases are not covered in UTs

## Part 2 notes

- So. It's the mask. Of course, there was a reason for the key to be a number.

### Floater strategy

We will loop through all the values from 0 to 2^_number of Xes in the mask_ and call them `i`.

`v` is the key to be modified. What we want is:
- a number, called `ones`, with all the ones for this instance of the floating mask marked as a one
- a number, called `zeros`, with all the zeros for this instance of the floating mask marked as a one
- a number, called `m.ones` with all the ones in the mask marked as a one

For example, if the mask is `1XX0`, we have 2 Xes, so `i` will go from `0` to `3`

| i  | ones | zeros | m.ones |
|----|------|-------|--------|
| 00 | 0000 | 0110  | 1000   |
| 01 | 0010 | 0100  | 1000   |
| 10 | 0100 | 0010  | 1000   |
| 11 | 0110 | 0000  | 1000   |

Then we just have to write ones where we have ones in `ones` and `m.ones`, and write zeros where we have ones in `zeros`.

See [the code](./docking/mask2.go) of the tricky function

[Read the story](./STORY.md)
