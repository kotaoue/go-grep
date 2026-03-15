# go-grep

[![Go](https://github.com/kotaoue/go-grep/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/kotaoue/go-grep/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/kotaoue/go-grep/branch/main/graph/badge.svg)](https://codecov.io/gh/kotaoue/go-grep)
[![Go Report Card](https://goreportcard.com/badge/github.com/kotaoue/go-grep)](https://goreportcard.com/report/github.com/kotaoue/go-grep)
[![License](https://img.shields.io/github/license/kotaoue/go-grep)](https://github.com/kotaoue/go-grep/blob/main/LICENSE)

Execute grep in Go.

## Installation

```bash
go get github.com/kotaoue/go-grep
```

## Example

```go
package main

import (
	"fmt"

	"github.com/kotaoue/go-grep/grep"
)

func main() {
	lines := []string{
		"apple",
		"banana",
		"apricot",
		"cherry",
	}

	results, err := grep.Find(lines, "ap")
	if err != nil {
		panic(err)
	}

	for _, line := range results {
		fmt.Println(line)
	}
}
```

**Output:**

```text
apple
apricot
```
