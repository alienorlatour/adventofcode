## Part 1 notes

- The docking package holds data structures for bit manipulation.
- The Parser does what its name hopefully suggests. I elected to use strings as the memory's keys, because parsing uint64 didn't seem worth it. After all, the keys are only that: keys to a map. No computation is needed on them.
- Corner/invalid cases are not covered in UTs
- 

[Read the story](./STORY.md)
