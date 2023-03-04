package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlres) AddProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.mgStore.ProductRepository().GetProduct(); err != nil {
			return fiber.ErrNotImplemented
		}
		return c.JSON(fiber.Map{
			"result": "success",
		})
	}
}
