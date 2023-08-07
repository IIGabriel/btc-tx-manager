package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	GetOne(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Custom(route string) func(ctx *fiber.Ctx) error
}
