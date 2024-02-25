## This is a sample set of Golang modules with goroutines and synchronization patterns

### How to start

- `go run main.go`
- `cat sample.txt | go run main.go` - calculate number of characters in the file

### Check directories for various concurency patterns

- [Pipeline](pipeline/): break into "stages" with a goroutine for each stage (communicate with channels)
- [Fan-in](fanin/): multiple input channels, single output channel
- [Close/cancelation](pipeline/): closing channel when end of data stream / end of work of a goroutine
- [Request Quit](cancellation/): dedicated "quit" channel. Listening for the channel and shuts down a goroutine
- [Context](ctx/): calling "quit" function cancells all operations using the Context
- [Generator](generator/): calculate "one at time" as per processed
- [all patterns: multi-threaded grep](mgrep/): an example how to utilize goroutines, channels, bufio, waitgroups

### Go CLI commands

- `go build` - builds & emits binary files
- `go build -race` - checks for concurrency problems
- `go mod tidy` - update dependencies
- `go test` - execute tests
- `go fmt` - format all source files

### Intricacies of specific modules

- module [pipeline](./pipeline/) has `image` package in use
- see an [additional information](https://golang.org/doc/articles/image_package.html) about it

### Kudos

- The examples came from the [Golang course, Zero To Mastery](https://academy.zerotomastery.io/courses/1600953/lectures/42962079)
