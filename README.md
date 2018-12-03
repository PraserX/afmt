# Advanced formatter

```
WORK IN PROGRESS
```

It is an extension of standard Golang formatter. The main feature of this extension is tree print.

## Example

```go
package main

import (
    "fmt"
    "net"
    "github.com/PraserX/afmt"
)

type TestStructure struct {
    Level0101 string
    Level0102 struct {
        Level0201 string
        Level0202 int
    }
    Level0103 bool
    Level0104 []string
}

func main() {
    str := TestStructure{}
    str.Level0101 = "Lorem ipsum dolor sit amet"
    str.Level0102.Level0201 = "Lorem ipsum dolor sit amet"
    str.Level0102.Level0202 = 10
    str.Level0103 = false
    str.Level0104 = []string{"Lorem", "ipsum", "dolor", "sit", "amet"}

    PrintTree(str)
}
```

The result of code above:

```
TestStructure:
├── Level0101: Lorem ipsum dolor sit amet
├── ■
│   ├── Level0201: Lorem ipsum dolor sit amet
│   └── Level0202: 10
├── Level0103: false
└── Level0104:
    ├── Lorem
    ├── ipsum
    ├── dolor
    ├── sit
    └── amet
```
