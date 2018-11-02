# Learning GoLang

You usually have to start with a [hello world](/01-hello/hello.go). But the best way to learn programming is with [TDD](/00-testing/first_test.go).

## Setup

First [install go](https://golang.org/doc/install) and then setup your [Go path](https://golang.org/doc/code.html#GOPATH).

Lets test the setup by creating a folder called `$GOPATH/src/learning`. Paste the below into a file called `hello.go`:

```go
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```

To execute use `go run hello.go`.

Go is also compiled into an executable. To do so run:

```bash
cd learning
go build hello.go // this creates the hello executable
./hello
```

The `package main` tells GoLang that this is an executable file and not a library. An executable file must also have a `func main()`.

The `fmt` is a package that handles input/output, in this case printing to stdout.

## Offline Go

If you have go installed, run:

```bash
godoc -http=:6060
```

Then open to http://127.0.0.1:6060/ in your browser.

## Packages

Grouping together code.

## Types

Strongly typed language.

`bool`, `string`, `int`, `float64`

## Variables

02-variables/variabls.go

Variable names. I can't beleive I have to mention this. Reading a lot of peoples interpretations of [pikestyle](http://doc.cat-v.org/bell_labs/pikestyle) reinforces my opinion of the human race. Oh yeah, you commute on the London tube every day for years and lets hear your opnion. Anyway, I digress. Pike clearly mentions 'clarity of expression'. The variable should describe exactly what it is. Our conventions for i in for loops is totally fine. Anyone who has programmed will know this one. That doesn't mean to say that all variables should be single letters does it? Code is written for humans. Beautiful code should read like well written prose.

## Numbers

[03-numbers](/03-numbers)

## Strings

[04-strings](/04-strings)

## Arrays and Slices

So Go does a lot of things for you. For this example we want to demonstrate arrays and slices. So lets have seperate files for this. Oh, so how do you import these files then? Aahhh, you don't have to do anything. Well, apart from change the run command to:

```bash
go run main.go array.go slice.go
go build main.go array.go slice.go
```

The run command needs to group all these files into an executable.

[05-array-slices](/05-array-slices)

Rarely use arrays directly as they are pretty inflexible. Use slices instead. Of course arrays have their uses.

You will see the reserved word 'make'. This function takes care of allocating memory for the slice in this example.

## Conditionals

[06-conditionals-iterators/statements.go](/06-conditionals-iterators)

## Functions

[07-functions/functions.go](/07-functions)

func [name]([[paramName/s,] type]) [[[name]return type], {(return types,)] {
  ...
}

## Structs

struct/struct.go

## Pointers

pointers/pointer.go

## Casting

## TDD

```bash
go test --coverage
```

testing/[coming soon]

## HTTP Server

http-server/server.go

## Time and Tickers

http-json/http-json.go

## Command Line

- [Go by example - flags](https://gobyexample.com/command-line-flags)

## References

- [Go Start](https://github.com/alco/gostart)
- [Glang writing unit tests](https://blog.alexellis.io/golang-writing-unit-tests/)
