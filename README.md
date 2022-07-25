# round-robin

A [go](https://go.dev/) implementation of the Round Robin balancing algorithm.

## Installation

To install round-robin package, you need to install Go and set your Go workspace first.

1. You first need Go installed, then you can use the below Go command to install round-robin.

```shell
go get -u github.com/stsmdt/round-robin
```

2. Import it in your code:

```go
import "github.com/stsmdt/round-robin"
```

## Example

```go
package main

import (
	"net/url"

	"github.com/stsmdt/round-robin"
)

func main() {
	roundRobin, _ := roundrobin.New(
		&url.URL{Host: "127.0.0.1"},
		&url.URL{Host: "127.0.0.2"},
		&url.URL{Host: "127.0.0.3"},
		&url.URL{Host: "127.0.0.4"},
	)

	roundRobin.Next() // {Host: "127.0.0.1"}
	roundRobin.Next() // {Host: "127.0.0.2"}
	roundRobin.Next() // {Host: "127.0.0.3"}
	roundRobin.Next() // {Host: "127.0.0.4"}
	roundRobin.Next() // {Host: "127.0.0.1"}
	roundRobin.Next() // {Host: "127.0.0.2"}
}
```

## License

This project is released under the [MIT License](LICENSE).
