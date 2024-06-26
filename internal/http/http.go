package http

import (
	"encoding/json"

	"go-key-value/internal/interfaces"
	"go-key-value/internal/models"

	"github.com/gofiber/fiber/v3"
)

func NewAuthRouting(r *fiber.App, a HTTPInterface) {
	r.Post("/v1/set", a.Set)
	r.Post("/v1/get", a.Get)
	r.Post("/v1/del", a.Del)

}

type HTTPInterface interface {
	Set(ctx fiber.Ctx) error
	Get(ctx fiber.Ctx) error
	Del(ctx fiber.Ctx) error
}

type Handler struct {
	redis interfaces.KeyValueRepositoryInterface
}

func NewHandler(redis interfaces.KeyValueRepositoryInterface) HTTPInterface {
	return &Handler{
		redis: redis,
	}
}

func (a *Handler) Set(ctx fiber.Ctx) error {
	request := &models.KeyValue{}

	if err := json.Unmarshal(ctx.Body(), request); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	responseData, err := a.redis.Set(ctx.Context(), request)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.JSON(responseData)
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return ctx.SendStatus(fiber.StatusOK)
}

func (a *Handler) Get(ctx fiber.Ctx) error {
	request := &models.KeyValue{}

	if err := json.Unmarshal(ctx.Body(), request); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	responseData, err := a.redis.Get(ctx.Context(), request)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.JSON(responseData)
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return ctx.SendStatus(fiber.StatusOK)
}

func (a *Handler) Del(ctx fiber.Ctx) error {
	request := &models.KeyValue{}

	if err := json.Unmarshal(ctx.Body(), request); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	responseData, err := a.redis.Del(ctx.Context(), request)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.JSON(responseData)
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return ctx.SendStatus(fiber.StatusOK)
}