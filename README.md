![](https://travis-ci.org/PraserX/afmt.svg?branch=master)

# Advanced formatter

It is an extension of standard Golang formatter. It contains (or it will contain) some advanced features for command line print. All features are described bellow.

## Features

- Line wrap (`PrintCol`)
- Tree print (`PrintTree`)

### PrintCol

The `PrintCol` function (and its derivations) allow wrap text based on integer specification. It means that you can set your stop mark and start print of text on the new line.

```go
afmt.PrintCol(80, "This is my longest text ever.")
```

### PrintTree

The TreePrinter expect any type as argument. It use interfaces, which are analyzed and printed. So you can use string, integer, array, structure and etc. The power of TreePrinter is that you can use any structure and TreePrinter prints it for you. There are many settings, which you can use.

The easiest way is to use function `PrintTree`: 

```go
afmt.PrintTree(myStructure);
```

On the other hand you can initialize TreePrinter with your own options:

```go
package main

import (
    "fmt"
    "os"
    "github.com/PraserX/afmt"
)

type MyStructure struct {
    Item1 string
    Item2 bool
    Item3 []string
}

func main() {
    var err error
    var result string

    str := MyStructure{}
    str.Item1 = "Lorem ipsum dolor sit amet"
    str.Item2 = false
    str.Item3 = []string{"Lorem", "ipsum"}

    tp := afmt.NewTreePrinter()

    if result, err = tp.Print(testValue); err != nil {
        fmt.Fprintf(os.Stderr, "Hmm something happened: %s", err.Error())
    }

    fmt.Printf(result)
}
```

So as a result you can see something like:

```
MyStructure:
├── Item1: Lorem ipsum dolor sit amet
├── Item2: false
└── Item3:
    ├── Lorem
    └── ipsum
```

## Planned features

- Text padding
- Progress bar

Is something missing? Create issue and describe your idea!

## License

This library is under MIT license.