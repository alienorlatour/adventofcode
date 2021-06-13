# Part 1

Yay! Trees and compilers!

After reading a great part of the Internet concerning compilers, AST, blog posts about grammars and a fair share of Wikipedia and Github, I think it's time to do it from scratch. The simplicity of the so-called language means over-engineering is a bad idea.

I started with this sort of interface:
```
// add.go
type addition struct {
    left  expression
    right expression
}

func (a addition) compute() int {
    return a.left.compute() + a.right.compute()
}

// scalar.go
type scalar struct {
	value int
}

func (a scalar) compute() int {
	return a.value
}
```

It was clever, really, but parsing using a turing machine means modifying the right-hand expression after setting the operator, which means I cannot keep the abstraction.


[Read the story](./STORY.md)
