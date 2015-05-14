[![GoDoc](https://godoc.org/github.com/ahamidi/go-mapreduce?status.svg)](https://godoc.org/github.com/ahamidi/go-mapreduce)

## go-mapreduce
Go Map/Reduce Package

## Functionality
The goal of `go-mapreduce` is to simplify the use of the common map/reduce pattern in Go. Essentially it takes the following 2 functions:

**Map**: Fan out function that typically retrieves data and applies some sort of transformation.

**Reduce**: Aggregation function that iterates over data emitted by the Map.

## TODO

- [x] Map Function
- [x] Reduce Function
- [ ] Report process stats
- [x] Docs
- [x] Tests
- [ ] Benchmarks

