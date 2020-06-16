# What is this?
Easy rounding

# Installation
```commandline
$ go get github.com/go-utils/round
```

# Usage
```go
package main

import (
    "fmt"

    "github.com/go-utils/round"
)

func main() {
    fmt.Println(round.Do(3.4444449))
    // Output: 4
}
```

# License
MIT
