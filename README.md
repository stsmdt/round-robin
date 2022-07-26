# round-robin

[![Run Tests](https://github.com/stsmdt/round-robin/actions/workflows/roundrobin.yml/badge.svg?branch=main)](https://github.com/stsmdt/round-robin/actions/workflows/roundrobin.yml)
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
import roundrobin "github.com/stsmdt/round-robin"
```

## Example

```go
package main

import (
	"net/url"

	roundrobin "github.com/stsmdt/round-robin"
)

func main() {
	rr, _ := roundrobin.New(
		[]url.URL{
			{Host: "127.0.0.1"},
			{Host: "127.0.0.2"},
			{Host: "127.0.0.3"},
			{Host: "127.0.0.4"},
			{Host: "127.0.0.5"},
		},
	)

	rr.Next() // {Host: "127.0.0.1"}
	rr.Next() // {Host: "127.0.0.2"}
	rr.Next() // {Host: "127.0.0.3"}
	rr.Next() // {Host: "127.0.0.4"}
	rr.Next() // {Host: "127.0.0.5"}
	rr.Next() // {Host: "127.0.0.1"}
}
```

## License

This project is released under the [MIT License](LICENSE).
