# Logur adapter for Uber's [Zap](https://github.com/uber-go/zap)

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/logur/adapter-zap/CI?style=flat-square)](https://github.com/logur/adapter-zap/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/logur/adapter-zap?style=flat-square)](https://codecov.io/gh/logur/adapter-zap)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/adapter/zap?style=flat-square)](https://goreportcard.com/report/logur.dev/adapter/zap)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/logur.dev/adapter/zap)


## Installation

```bash
go get logur.dev/adapter/zap
```


## Usage

```go
package main

import (
	"go.uber.org/zap"
	zapadapter "logur.dev/adapter/zap"
)

func main() {
	logger := zapadapter.New(zap.NewNop())
}
```


## Development

When all coding and testing is done, please run the test suite:

```bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
