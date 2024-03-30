package auth

import (
	"encoding/json"

	"redis/internal/redisRepository"
	"redis/internal/models"

	"github.com/gofiber/fiber/v3"
)

func NewAuthRouting(r *fiber.App, a *AuthHandler) {
	r.Post("/v1/set", a.Set)
	r.Get("/v1/get", a.Get)
}

type AuthHandlerInterface interface {
	Set(ctx fiber.Ctx) error
	Get(ctx fiber.Ctx) error
}

type AuthHandler struct {
	redis *redisrepository.RedisRepository
}

func NewAuthHandler(redis *redisrepository.RedisRepository) *AuthHandler {
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

	response, err := json.Marshal(responseData)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	if err = ctx.Send(response); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

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

	response, err := json.Marshal(responseData)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	if err = ctx.Send(response); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}