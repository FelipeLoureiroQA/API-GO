package main

import (
	"fmt"

	"github.com/FelipeLoureiroQA/API-GO/config"
	"github.com/FelipeLoureiroQA/API-GO/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Err(fmt.Sprintf("Error initializing config: %v", err))

		return
	}

	router.Initialize()
}
