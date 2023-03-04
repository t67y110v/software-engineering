package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlers) AddProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {

		if err := h.mgStore.ProductRepository().GetProduct(); err != nil {
			return fiber.ErrNotImplemented
		}

		return c.JSON(fiber.Map{
			"result": "success",
		})
	}
}

func (h *Handlers) GetAllProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		products, err := h.mgStore.ProductRepository().GetAllProducts()
		if err != nil {
			return err
		}
		return c.JSON(products)
	}
}
