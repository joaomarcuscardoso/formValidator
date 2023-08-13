package main

import (
	"github.com/joaomarcuscardoso/formValidator/internal/config"
	"github.com/joaomarcuscardoso/formValidator/internal/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	// Initalize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// Initialize router
	router.Initialize()
}
