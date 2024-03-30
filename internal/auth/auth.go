package auth

import (
	"encoding/json"

	"redis/internal/redisRepository"
	"redis/internal/models"

	"github.com/gofiber/fiber/v3"
)

func NewAuthRouting(r *fiber.App, a AuthHandlerInterface) {
	r.Post("/v1/set", a.Set)
	r.Get("/v1/get", a.Get)
}

type AuthHandlerInterface interface {
	Set(ctx fiber.Ctx) error
	Get(ctx fiber.Ctx) error
}

type AuthHandler struct {
	redis redisrepository.RedisInterface
}

func NewAuthHandler(redis redisrepository.RedisInterface) AuthHandlerInterface {
	return &AuthHandler{
		redis: redis,
	}
}

func (a *AuthHandler) Set(ctx fiber.Ctx) error {
	request := &models.Redis{}

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

func (a *AuthHandler) Get(ctx fiber.Ctx) error {
	request := &models.Redis{}

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