# round-robin

[![Run Tests](https://github.com/stsmdt/round-robin/actions/workflows/roundrobin.yml/badge.svg?branch=main)](https://github.com/stsmdt/round-robin/actions/workflows/roundrobin.yml)
[![codecov](https://codecov.io/gh/stsmdt/round-robin/branch/main/graph/badge.svg?token=CY1OHFI4YU)](https://codecov.io/gh/stsmdt/round-robin)
[![Go Report Card](https://goreportcard.com/badge/github.com/stsmdt/round-robin)](https://goreportcard.com/report/github.com/stsmdt/round-robin)
[![GoDoc](https://pkg.go.dev/badge/github.com/stsmdt/round-robin?status.svg)](https://pkg.go.dev/github.com/stsmdt/round-robin?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/stsmdt/round-robin/-/badge.svg)](https://sourcegraph.com/github.com/stsmdt/round-robin?badge)
[![Release](https://img.shields.io/github/release/stsmdt/round-robin.svg?style=flat-square)](https://github.com/stsmdt/round-robin/releases)

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
