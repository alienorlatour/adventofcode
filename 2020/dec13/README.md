#### Part 2 notes

My first attempt looked like this:

```go
func (b busRule) EarliestRuleValidation() (int64, error) {
	firstBus, ok := b.rule[0]
	if !ok {
		return 0, fmt.Errorf("invalid input, starting with x is not supported")
	}
	// try each time the first bus departs
	for time := int64(0); ; time += int64(firstBus) {
		if b.isValid(time) {
			return time, nil
		}
	}
}
```
This takes an awful lot of time.

I then tried with the largest ID as a step, because finding a more complex algorithm would require more time than I have. It still takes hours, but I can wait.

I would like to benchmark this version along with the use of the most magic GO feature: paralellisation.

Meanwhile, for the sake of future users, I added the `startAt` parameter, to quicken things up a bit...


[Read the story](./STORYhttps://github.com/login?return_to=https%3A%2F%2Fgithub.com%2Fablqk%2Fadventofcode%2Fpull%2Fnew%2Ffeature%2Fstorytelling.md)
