package main

import (
	"product-management-service/internal/app"
	"product-management-service/internal/config"
)

func main() {
	cfg := config.NewConfig()
	fiberApp := config.NewFiber(cfg)
	app.Container(fiberApp, cfg)
	app.RunWithGracefulShutdown(fiberApp, cfg)
}
