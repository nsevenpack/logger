# logger

[![Go Reference](https://pkg.go.dev/badge/github.com/nsevenpack/logger.svg)](https://pkg.go.dev/github.com/nsevenpack/logger)
[![Version](https://img.shields.io/github/v/tag/nsevenpack/logger?label=version&sort=semver)](https://github.com/nsevenpack/logger/releases)
[![CI](https://github.com/nsevenpack/logger/actions/workflows/release.yml/badge.svg)](https://github.com/nsevenpack/logger/actions/workflows/release.yml)
[![License](https://img.shields.io/github/license/nsevenpack/logger)](https://github.com/nsevenpack/logger/blob/main/LICENSE)


Petit package de log Go simple.

## Installation

```bash
# installe la derniere version
go get github.com/nsevenpack/logger@latest

# liste les versions disponibles
go list -m -versions github.com/nsevenpack/logger

# installe une version précise
go get github.com/nsevenpack/logger@v1.0.0
```
## Utilisation

```golang
package main

import "github.com/nsevenpack/logger"

func main() {
	logger.Info("Hello!")
}
```