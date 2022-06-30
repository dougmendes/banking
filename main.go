package main

import (
	"banking/app"
	"banking/logger"
)

func main() {

	logger.Info("Starting application")
	app.Start()
}
