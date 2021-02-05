# Part 1

## Rules

- According to a quick look into the input, all rules are declared as two ranges of ints.

# Part 2

We'll need to match each rule with a possible order. The first solution that comes to mind is to keep a matrix of all possibilities and remove each impossible position of a rule when a ticket is parsed. Empirically, it seems rather heavy on memory.

A second idea would be to use the first ticket to create that matrix, so that is starts smaller. This optimisation will stay an idea.

In both cases, the rule will have to keep a list of its possible positions and clear it until it has only one element.

After trying a few different approaches, I conclude that bitwise is better. I read some stuff about bitwise computation a few days ago and now I want to place it everywhere. This would never happen with graphs and trees.

Assumption: there is a solution. In the case where there is no solution, one would be left with an infinite loop.

[Read the story](./STORY.md)
