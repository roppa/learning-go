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

The `package main` tells GoLang that this is an executable file and not a library. An executable file must also have a `func main()`. If it is named something other than `main` that would be a library.

The `fmt` is a package that handles input/output, in this case printing to stdout.

## Offline Go

If you have go installed, run:

```bash
godoc -http=:6060
```

Then open to http://127.0.0.1:6060/ in your browser.

## Packages/Libraries

A package, or library, groups together code which is utilised by another program.

## Types

Strongly typed language.

`bool`, `string`, `int`, `float64`

Go word size is the memory address size. Based on your machine, 32bit, 64bit etc. An integer is based on your word size, so on a 64bit machine it will be 64 bits long, 32bits on a 32 bit machine.

Go zero values:

- nil
- 0
- false
- “”

## Testing and TDD

Testing, like a lot of languages, is included in GoLang.

[00-testing](/00-testing)

To run a test:

```bash
go test --cover
```

## Variables

[02-variables](02-variables/variabls.go)

Variable names. I can't beleive I have to mention this. Reading a lot of peoples interpretations of [pikestyle](http://doc.cat-v.org/bell_labs/pikestyle) reinforces my opinion of the human race. Oh yeah, you commute on the London tube every day for years and lets hear your opnion. Anyway, I digress. Pike clearly mentions 'clarity of expression'. The variable should describe exactly what it is. Our conventions for i in for loops is totally fine. Anyone who has programmed will know this one. That doesn't mean to say that all variables should be single letters does it? Code is written for humans. Beautiful code should read like well written prose.

You can do multiple assignment:

```golang
a, b, c := 10, 100, 1000
fmt.Println(a, b, c)
```

## Numbers

[03-numbers](/03-numbers)

## Strings

[04-strings](/04-strings)

## Arrays and Slices

[05-array-slices](/05-array-slices)

So Go does a lot of things for you. For this example we want to demonstrate arrays and slices. So lets have seperate files for this. Oh, so how do you import these files then? Aahhh, you don't have to do anything. Well, apart from change the run command to:

```bash
go run main.go array.go slice.go
go build main.go array.go slice.go
```

The run command needs to group all these files into an executable.

Rarely use arrays directly as they are pretty inflexible. Use slices instead. Of course arrays have their uses.

You will see the reserved word 'make'. This function takes care of allocating memory for the slice in this example.

`make` does a few things, one is creating a channel, takes care of allocating memory for a slice.

## Conditionals

[06-conditionals-iterators/statements.go](/06-conditionals-iterators)

## Functions

[07-functions/functions.go](/07-functions)

func [name]([[paramName/s,] type]) [[[name]return type], {(return types,)] {
...
}

Functions beginning with a capital letter are exported. Fat.Println etc, notice the capital P.

## Structs

A struct is a custom type, which is an aggregate of other types.

In structs, always have attributes listed from shortest length of the data type. int8, int16 etc as go pads out the structure with memory allocation based on the first longest?

[struct/struct.go](struct/struct.go)

## Pointers

Go is pass by value. To pass by reference you need to use pointers in the calling function parameter (\*) and use 'address of' in the callee function parameters and body (&) &[variable name] i.e. &count.

- `&` - 'the address of'
- `*` - 'the value of'

```go
  hello := "Hello"
  var helloPointer *string = &hello
  fmt.Println(&hello) // 0xc0000101e0
  fmt.Println(*&hello) // Hello
  fmt.Println(*helloPointer) // Hello
```

[pointers/pointer.go](./pointers)

## Go routines

All Go you write gets run on a goroutine - an OS thread.

## Channels

Channels are pipes that connect concurrent goroutines.

## HTTP Server

http-server/server.go

## Pipes

- [Golang pipes](https://golang.org/src/io/pipe.go)
- [Zup Zup pipes](https://zupzup.org/io-pipe-go/)

## Time and Tickers

http-json/http-json.go

## Command Line

- [Go by example - flags](https://gobyexample.com/command-line-flags)

## Race Conditions

- [GoLang data race and how to fix it](https://www.sohamkamani.com/blog/2018/02/18/golang-data-race-and-how-to-fix-it/)

## References

- [Go Start](https://github.com/alco/gostart)
- [Glang writing unit tests](https://blog.alexellis.io/golang-writing-unit-tests/)
- [How to write Go code](https://golang.org/doc/code.html).
- [Go by example](https://gobyexample.com/)
- [Learn go with tests](https://github.com/quii/learn-go-with-tests)
- [GoLang tour](https://tour.golang.org/list)
- [Awesome GoLang information from rakyll.org!](https://rakyll.org/)
- [TDD](https://leanpub.com/golang-tdd/read)
- [Web apps with GoLang](https://astaxie.gitbooks.io/build-web-application-with-golang/en/)
- [GoLang talks](https://talks.golang.org/2012/10things.slide#8)
- [Thinking in Go](https://hackmongo.com/post/thinking-in-go/)
- [Ultimate Go programming](https://www.safaribooksonline.com/videos/ultimate-go-programming/9780134757476/9780134757476-ugpg_02_02_03_02)
- [Programming guide](https://programming.guide/go/three-dots-ellipsis.html)
- [research.swtch.com](https://research.swtch.com/)
- [golangbot.com](https://golangbot.com)
