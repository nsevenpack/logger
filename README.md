# logger

Petit package de log Go simple.

## Installation

```bash
go get github.com/ton-org/logger@latest
```
## Utilisation

```go
package main

import "github.com/ton-org/logger"

func main() {
	logger.Info("Hello!")
}
```

## Dev publication  

```bash
git commit -m "feat: add new feature" # ==> 0.1.0
git commit -m "fix: fix bug" # ==> 0.0.1
git commit -m "feat!: update dependencies" # ==> 1.0.0

