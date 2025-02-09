# DSA - Datastructures and algorithms

Project to learn them all. Maybe. At some point. Probably not :D

For now everything written in go, because it is simple, but not javascript "simple".

---

## Running algorithms

Running datastructures and algorithms is done via tests.
There is no main package in this repository.

---

## Testing

Every datastructure and algorithm has their own test files.

Test individually by cd'ing into their directory and `go test`

Or test everything from the root with `go test ./...`

Get test coverage with `go test -cover ./...`

## Testing additions: Sugar

This repository also includes the sugar library, which adds some coloring, time and memory metrics to testing.
Needs the performance and color library alongside.

All can be found in the DSA/util directory.