package app

import (
	"product-management-service/internal/config"
	"product-management-service/internal/delivery/http/handler"
	"product-management-service/internal/delivery/http/middleware"
	"product-management-service/internal/delivery/http/router"
	"product-management-service/internal/repository"
	"product-management-service/internal/usecase/product"

	"github.com/gofiber/fiber/v2"
)

func Container(fiberApp *fiber.App, cfg *config.Config) {
	// init infrastruture
	logger := config.NewLogger(cfg)
	db := config.NewDatabase(cfg, logger)
	validate := config.NewValidator()
	redisClient := config.NewRedis(cfg)

	//init  repositories
	productRepo := repository.NewProductRepository()
	redisRepo := repository.NewRedisRepository(redisClient)

	//init  usecases
	productUsecase := product.NewProductUsecase(db, productRepo, redisRepo, validate)

	//init  handler
	pingHandler := handler.NewPingHandler()
	productHandler := handler.NewProductHandler(productUsecase, logger)

	//init  middleware
	loggingMiddleware := middleware.HandleReqLogging(logger)
	recoveryMiddleware := middleware.HandleRecoveryPanic()

	route := &router.Route{
		App:               fiberApp,
		RecoverMiddleware: recoveryMiddleware,
		LogMiddleware:     loggingMiddleware,
		PingHandler:       pingHandler,
		ProductHandler:    productHandler,
	}
	route.RegisterRoutes()
}
