![](https://github.com/ablqk/adventofcode/workflows/validation/badge.svg)

# Advent of Code

This is my take on the [2020 Advent of Code](https://adventofcode.com/2020) using the GO language.

Each day of December is found under dec followed by the two-digit number of the day. A README contains a copy of the subject, the main file is also named after the day.

## Quickstart

### Prerequisites

You will need a [functionning version of go](https://golang.org/doc/install). I'm using the latest to this day, which is 1.15.

The repo can be used without the make command, but it is highly recommanded.

### Run

1. Checkout the project, then download vendors:
    ```
    make init
    ```
1. Run, giving the number of the day as parameter:
    ```
    make run -- 1
    ```
    or
    ```
    make build
    ./adventofcode.out -day 1
    ```

You can also run unit tests with coverage:
```
make test
```

# What could be improved

- Most days are covered with a very minimal and nominal test case, using the example given by the subject. Corner cases should be covered.
- I took the liberty of a few [0] without checking the length of the slice. These occasions are systematically accompanied by a `// todo`. Given the importance of the code and the ungodly hour of the night, I can only hope that some day the will is there to correct this. Don't do this at home.
- The methods of the fileread package should probably ignore empty lines (December 4: no). A lot of shortcuts are made here, as the inputs are totally under control. Error management could use some consolidation.
- Some doors' code tends to reside in the func given as parameter to readAndApply. I would like to find a way to pull these out for unit testing and readbility.
- I am particularly not proud of the quality found in Dec 7th. The comments look like the Dune planet and the naming is Lewis Carollesque.

