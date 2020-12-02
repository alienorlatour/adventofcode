![](https://github.com/ablqk/adventofcode/workflows/validation/badge.svg)

# Advent of Code

This is my take on the [2020 Advent of Code](https://adventofcode.com/2020) using the GO language.

Each day of December is found under dec followed by the two-digit number of the day. A README contains a copy of the subject, the main file is also named after the day.

## Quickstart

1. Checkout the project, then download vendors:
    ```
    make init
    ```
1. Run, giving the number of the day as parameter:
    ```
    make run 1
    ```
    or
    ```
    make build
    ./adventofcode.out -day 1
    ```

You can also run unit tests:
```
make test
```

# What could be improved

- Most days are covered with a very minimal and nominal test case, using the example given by the subject. Corner cases should be covered.
- I took the liberty of a few [0] without checking the length of the slice. These occasions are systematically accompanied by a `// todo`. Given the importance of the code and the ungodly hour of the night, I can only hope that some day the will is there to correct this. Don't do this at home.
- The methods of the fileread package should probably ignore empty lines. A lot of shourtcuts are made here, as the inputs are totally under control. Error management could use some consolidation.
