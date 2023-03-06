package handlers

import (
	"bytes"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetCart() fiber.Handler {
	return func(c *fiber.Ctx) error {
		params := c.Params("user_id")

		p, err := h.mgStore.ProductRepository().GetCart(params)
		if err != nil {
			return err
		}

		return c.JSON(p)
	}
}

func (h *Handlers) AddToCart() fiber.Handler {

	return func(c *fiber.Ctx) error {

		type request struct {
			UserId  string `json:"user_id"`
			Product string `json:"product_name"`
		}

		req := &request{}

		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}

		if err := h.mgStore.ProductRepository().AddToCart(req.UserId, req.Product); err != nil {
			return err
		}

		cart, err := h.mgStore.ProductRepository().GetCart(req.UserId)
		if err != nil {
			return err
		}

		return c.JSON(cart)

	}
}

func (h *Handlers) DeleteFromCart() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			UserId  string `json:"user_id"`
			Product string `json:"product_name"`
		}

		req := &request{}

		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}

		if err := h.mgStore.ProductRepository().DeleteFromCart(req.UserId, req.Product); err != nil {
			return err
		}

		return nil
	}
}

func (h *Handlers) ClearCart() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			UserId string `json:"user_id"`
		}

		req := &request{}

		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}

		if err := h.mgStore.ProductRepository().ClearCart(req.UserId); err != nil {
			return err
		}

		return nil
	}
}
