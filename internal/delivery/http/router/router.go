package router

import (
	"product-management-service/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Route struct {
	App               *fiber.App
	RecoverMiddleware fiber.Handler
	LogMiddleware     fiber.Handler
	PingHandler       *handler.PingHandler
	ProductHandler    *handler.ProductHandler
}

func (r *Route) RegisterRoutes() {
	r.App.Use(r.RecoverMiddleware)
	r.App.Use(r.LogMiddleware)
	r.App.Use(cors.New())
	r.App.Get("/ping", r.PingHandler.Ping)

	v1Group := r.App.Group("/api/v1")

	productGroup := v1Group.Group("/products")
	productGroup.Post("/", r.ProductHandler.Create)
	productGroup.Patch("/:id", r.ProductHandler.Update)
	productGroup.Delete("/:id", r.ProductHandler.Delete)
	productGroup.Get("/:id", r.ProductHandler.GetDetail)

}
