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
git commit -m "fix: improve error output"
# => 1.0.1

git commit -m "feat: support custom log level"
# => 1.1.0

git commit -m "feat!: change default log level behavior"
# => 2.0.0
```

