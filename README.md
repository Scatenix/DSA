# DSA - Datastructures and Algorithms

Project to learn & implement datastructures and algorithms with Test-Driven-Development and as close to 100% test coverage as making sense and is possible.

Written in go, because it is simple and has type safety.

## Tech Stack

- languages: go
- libraries: None (aiming to implement everything from scratch)

## Structure

- `algorithms/` — implementations of various algorithms, grouped by category  
- `datastructures/` — implementations of fundamental data structures  
- `playground/` — application entry point with temporary code for quick experiments and manual testing  
- `util/` — utility functions, including a custom test helper package called `sugar`

## Running the Algorithms

Running and testing the datastructures and algorithms is done via unit tests.
The `playground` package in this repository is purley to quickly test things with temporary code.

Tests are always located next to the DSA implementation in the same directory.

## Testing

The goal of this repo, next to learning DSA, is to do Test-Driven-Development.

Every datastructure has their own test files.

Algorithms are grouped into their category (like sorting) and then run with parametrized unit tests if possible.

### How to test

Test individually by cd'ing into their directory and `go test`

Test everything from the root with `go test ./...`

Get test coverage with `go test -cover ./...`

## Testing additions: Sugar

This repository also includes a package called 'sugar' (sugar.go) written by me, which adds some coloring, time and memory metrics to testing.
Needs the performance and color package alongside. (Currently not yet in any separate repository)

Can be found in the DSA/util directory.
