package main

import (
	"github.com/magrininicolas/gopportunities-tutorial/config"
	"github.com/magrininicolas/gopportunities-tutorial/router"
)

func main() {
	logger := *config.GetLogger("main")
	if err := config.Init(); err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
