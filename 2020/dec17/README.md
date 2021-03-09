# Part 1

These things are rather hard to debug.

I chose a different approach from the last life-game door: while looping over the cells, each one checks how many neighbours it has rather than adding itself to its neighbours' counts. It is probably far less efficient, but way easier to test as it feels more natural.

# Part 2

OK. Great. 4 dimensions. Should I support N dimensions or just wrap `for` loops around my loops? Declare the number of dimensions as a const and work from there?

It seems at this moment that fast is the priority.


[Read the story](./STORY.md)
