package main

import "github.com/nsevenpack/logger/logger"

func main() {
	logger.Init()
	defer logger.Close()

	succes := "This is a success message" 
	info := "This is a info message"
	warning := "This is a warning message"
	errorMessage := "This is a error message"
	fatal := "This is a error fatal message"

	logger.S(succes)
	logger.I(info)
	logger.W(warning)
	logger.E(errorMessage)

	logger.Sf("Success message %v", succes)
	logger.If("Info message %v", info)
	logger.Wf("Warning message %v", warning)
	logger.Ef("Error message %v", errorMessage)
	logger.Ff("Fatal message %v", fatal)
}