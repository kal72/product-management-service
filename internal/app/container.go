package app

import (
	"product-management-service/internal/config"
	"product-management-service/internal/delivery/http/handler"
	"product-management-service/internal/delivery/http/middleware"
	"product-management-service/internal/delivery/http/router"
	"product-management-service/internal/usecase/auth"

	"github.com/gofiber/fiber/v2"
)

func Container(fiberApp *fiber.App, cfg *config.Config) {
	// init infrastruture
	logger := config.NewLogger(cfg)
	// db := config.NewDatabase(cfg, logger)
	// validate := config.NewValidator()

	//init  repositories

	//init  usecases
	authUsecase := auth.NewAuthUsecase(cfg)

	//init  handler
	pingHandler := handler.NewPingHandler()

	//init  middleware
	loggingMiddleware := middleware.HandleReqLogging(logger)
	authMiddleware := middleware.HandleAuth(authUsecase)

	route := &router.Route{
		App:            fiberApp,
		LogMiddleware:  loggingMiddleware,
		AuthMiddleware: authMiddleware,
		PingHandler:    pingHandler,
	}
	route.RegisterRoutes()
}
