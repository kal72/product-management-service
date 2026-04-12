package handler

import (
	"product-management-service/internal/model"
	"product-management-service/internal/usecase/product"
	"product-management-service/internal/utils/errorhandler"
	"product-management-service/internal/utils/logger"
	"product-management-service/internal/utils/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	usecase product.ProductUsecaseContract
	log     *logger.Logger
}

func NewProductHandler(usecase product.ProductUsecaseContract, log *logger.Logger) *ProductHandler {
	return &ProductHandler{
		usecase: usecase,
		log:     log,
	}
}

func (h *ProductHandler) Create(ctx *fiber.Ctx) error {
	var req model.CreateProductRequest
	var errData *model.ErrorData
	var resp any

	defer func() {
		var err error
		if errData != nil {
			err = errData.Error
			if resp == nil {
				resp = errData
			}
		}
		h.log.LogEvent(ctx.UserContext(), ctx.Response().StatusCode(), err, req, resp)
	}()

	if err := ctx.BodyParser(&req); err != nil {
		errData = errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	product, errRet := h.usecase.Create(ctx.UserContext(), &req)
	if errRet != nil {
		errData = errRet
		return response.ResponseError(ctx, errData)
	}

	resp = product
	return response.ResponseSuccess(ctx, product)
}

func (h *ProductHandler) Update(ctx *fiber.Ctx) error {
	var req model.UpdateProductRequest
	var errData *model.ErrorData

	defer func() {
		var err error
		if errData != nil {
			err = errData.Error
		}

		resp := string(ctx.Response().Body())
		h.log.LogEvent(ctx.UserContext(), ctx.Response().StatusCode(), err, req, resp)
	}()

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errData = errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	if err := ctx.BodyParser(&req); err != nil {
		errData = errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	if errData = h.usecase.Update(ctx.UserContext(), id, &req); errData != nil {
		return response.ResponseError(ctx, errData)
	}

	return response.ResponseSuccess(ctx, nil)
}

func (h *ProductHandler) Delete(ctx *fiber.Ctx) error {
	var errData *model.ErrorData
	var resp interface{}

	defer func() {
		var err error
		if errData != nil {
			err = errData.Error
			if resp == nil {
				resp = errData
			}
		}
		h.log.LogEvent(ctx.UserContext(), ctx.Response().StatusCode(), err, nil, resp)
	}()

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errData = errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	if errRet := h.usecase.Delete(ctx.UserContext(), id); errRet != nil {
		errData = errRet
		return response.ResponseError(ctx, errData)
	}

	resp = "success"
	return response.ResponseSuccess(ctx, nil)
}

func (h *ProductHandler) GetDetail(ctx *fiber.Ctx) error {
	var errData *model.ErrorData

	defer func() {
		var err error
		if errData != nil {
			err = errData.Error
		}

		resp := string(ctx.Response().Body())
		h.log.LogEvent(ctx.UserContext(), ctx.Response().StatusCode(), err, nil, resp)
	}()

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errData = errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	product, errRet := h.usecase.GetDetailByID(ctx.UserContext(), id)
	if errRet != nil {
		errData = errRet
		return response.ResponseError(ctx, errData)
	}

	return response.ResponseSuccess(ctx, product)
}

func (h *ProductHandler) List(ctx *fiber.Ctx) error {
	var req model.ProductFilter
	var errData *model.ErrorData

	defer func() {
		var err error
		if errData != nil {
			err = errData.Error
		}

		resp := string(ctx.Response().Body())
		h.log.LogEvent(ctx.UserContext(), ctx.Response().StatusCode(), err, req, resp)
	}()

	if err := ctx.QueryParser(&req); err != nil {
		errData = errorhandler.ErrorInvalidRequest(err)
		return response.ResponseError(ctx, errData)
	}

	products, pages, errRet := h.usecase.List(ctx.UserContext(), &req)
	if errRet != nil {
		errData = errRet
		return response.ResponseError(ctx, errData)
	}

	return response.ResponseSuccessWithPagination(ctx, products, pages)
}
