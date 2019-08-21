# Logur adapter for Uber's [Zap](https://github.com/uber-go/zap)

[![CircleCI](https://circleci.com/gh/logur/adapter-zap.svg?style=svg)](https://circleci.com/gh/logur/adapter-zap)
[![Coverage](https://gocover.io/_badge/logur.dev/adapter/zap)](https://gocover.io/logur.dev/adapter/zap)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/adapter/zap?style=flat-square)](https://goreportcard.com/report/logur.dev/adapter/zap)
[![GolangCI](https://golangci.com/badges/github.com/logur/adapter-zap.svg)](https://golangci.com/r/github.com/logur/adapter-zap)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)](https://github.com/logur/adapter-zap)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/logur.dev/adapter/zap)


## Installation

```bash
go get logur.dev/adapter/zap
```


## Usage

```go
package main

import (
	zapadapter "logur.dev/adapter/zap"
)

func main() {
	logger := zapadapter.New(/*logger*/)
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
