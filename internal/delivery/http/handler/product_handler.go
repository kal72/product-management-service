package handler

import (
	"product-management-service/internal/model"
	"product-management-service/internal/usecase/product"
	"product-management-service/internal/utils/errorhandler"
	"product-management-service/internal/utils/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	usecase product.ProductUsecaseContract
}

func NewProductHandler(usecase product.ProductUsecaseContract) *ProductHandler {
	return &ProductHandler{
		usecase: usecase,
	}
}

func (h *ProductHandler) Create(ctx *fiber.Ctx) error {
	var req model.CreateProductRequest
	if err := ctx.BodyParser(&req); err != nil {
		errData := errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	product, errData := h.usecase.Create(ctx.UserContext(), &req)
	if errData != nil {
		return response.ResponseError(ctx, errData)
	}

	return response.ResponseSuccess(ctx, product)
}

func (h *ProductHandler) Update(ctx *fiber.Ctx) error {
	var req model.UpdateProductRequest

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errData := errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	if err := ctx.BodyParser(&req); err != nil {
		errData := errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	errData := h.usecase.Update(ctx.UserContext(), id, &req)
	if errData != nil {
		return response.ResponseError(ctx, errData)
	}

	return response.ResponseSuccess(ctx, nil)
}

func (h *ProductHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errData := errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	errData := h.usecase.Delete(ctx.UserContext(), id)
	if errData != nil {
		return response.ResponseError(ctx, errData)
	}

	return response.ResponseSuccess(ctx, nil)
}

func (h *ProductHandler) GetDetail(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errData := errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	product, errData := h.usecase.GetDetailByID(ctx.UserContext(), id)
	if errData != nil {
		return response.ResponseError(ctx, errData)
	}

	return response.ResponseSuccess(ctx, product)
}

func (h *ProductHandler) List(ctx *fiber.Ctx) error {
	var req model.ProductFilter
	if err := ctx.QueryParser(&req); err != nil {
		errData := errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	products, pages, errData := h.usecase.List(ctx.UserContext(), &req)
	if errData != nil {
		return response.ResponseError(ctx, errData)
	}

	return response.ResponseSuccessWithPagination(ctx, products, pages)
}
