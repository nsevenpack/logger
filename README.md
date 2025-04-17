# logger

[![Go Reference](https://pkg.go.dev/badge/github.com/nsevenpack/logger.svg)](https://pkg.go.dev/github.com/nsevenpack/logger)
[![Version](https://img.shields.io/github/v/tag/nsevenpack/logger?label=version&sort=semver)](https://github.com/nsevenpack/logger/releases)
[![CI](https://github.com/nsevenpack/logger/actions/workflows/release.yml/badge.svg)](https://github.com/nsevenpack/logger/actions/workflows/release.yml)
[![License](https://img.shields.io/github/license/nsevenpack/logger)](https://github.com/nsevenpack/logger/blob/main/LICENSE)


Petit package de log Go simple.

## Installation

```bash
# installe la derniere version 1.x.x
go get github.com/nsevenpack/logger@latest
# installe la derniere version 2.x.x
go get github.com/nsevenpack/logger/v2@latest

# liste les versions disponibles pour 1.x.x
go list -m -versions github.com/nsevenpack/logger
# liste les versions disponibles pour 2.x.x
go list -m -versions github.com/nsevenpack/logger/v2

# installe une version précise
go get github.com/nsevenpack/logger@v1.1.0
# ou par version majeure
go get github.com/nsevenpack/logger/v2@v2.0.1
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
	// initialisé le logger, la creation du fichier de log et du dossier
	logger.Init()
	// defer pour close à la fin du programme
	defer logger.Close()

	// log success
	logger.S("Hello!")
	// log error
	logger.E("Hello!")
	// log warning
	logger.W("Hello!")
	// log info
	logger.I("Hello!")
	// log fatal (attention celui ci arrete le programme)
	logger.F("Hello!")

	// log success avec une variable
	logger.Sf("Hello %s", "world")
	// log error avec une variable
	logger.Ef("Hello %s", "world")
	// log warning avec une variable
	logger.Wf("Hello %s", "world")
	// log info avec une variable
	logger.If("Hello %s", "world")
	// log fatal avec une variable (attention celui ci arrete le programme)
	logger.Ff("Hello %s", "world")
}
```