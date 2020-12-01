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
