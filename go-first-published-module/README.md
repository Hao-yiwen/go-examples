# go-first-published-module

A simple Go module example for learning how to publish Go packages.

## Installation

```bash
go get github.com/Hao-yiwen/go-examples/go-first-published-module
```

## Usage

```go
package main

import (
    "fmt"
    hello "github.com/Hao-yiwen/go-examples/go-first-published-module"
)

func main() {
    fmt.Println(hello.SayHello())
}
```

## API

### SayHello

```go
func SayHello() string
```

Returns a greeting message "Hello, World!".

## License

MIT
