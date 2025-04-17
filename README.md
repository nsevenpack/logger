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

## Fonctionnalités
- Le logger creer un dossier tmp/log/dev par defaut si vous avez une variable d'environnement APP_ENV=prod ou APP_ENV=test etc ...  
alors le dossier dev sera remplacé par le nom de l'environnement.  
- Le logger creer un fichier de log par jour.  
- Jounalisation dans le fichier créé et dans le terminal.  

## Utilisation

```golang
package main

import "github.com/nsevenpack/logger"

func main() {
	logger.Info("Hello!")
}
```