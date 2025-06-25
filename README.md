kmpr is a tool to compare two values with the same type.
It is tested to compare slice, struct, function, and interface.

## Installation

Before installing, [download and install Go](https://go.dev/doc/install). Go 1.24.3 or higher required.
Installation is done using `go get` command.

```bash
go get -u github.com/lutffmn/kmpr
```

Or using the alternative way.
Import the module directly in source code.

```go
import "github.com/lutffmn/kmpr"
```

Then install it using `go mod` tool.

```bash
go mod tidy
```

## Features

- Slice Comparison
- Struct Comparison
- Function Comparison

## Examples

- Slice Comparison

```go

package main

import (
  "fmt"

  "github.com/lutffmn/kmpr"
)

func main(){
  fruits:= []string{"apple", "banana", "mango"}
  numbers:= []int{1, 2, 3, 4}

  if kmpr.Do(fruits, numbers){
    fmt.Println("These are the same slices!")
  }else{
    fmt.Println("These are two different slices!")
  }
}
```

## Issues

If you discover an issue, please inform me by submit an issue.
