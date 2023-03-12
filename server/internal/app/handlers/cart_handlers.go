package handlers

import (
	"bytes"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// @Summary Gets users cart
// @Description Getting users card by user_id
// @Tags         Cart
//
//	@Accept       json
//
// @Param        user_id   path      string  true  "User_id"
// @Produce json
// @Success 200 {object} responses.AllProducts
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /cart/get/{user_id} [get]
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

// @Summary Add to cart
// @Description Add product to users cart
// @Tags         Cart
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.AddToCart true  "Add to cart"
// @Success 200 {object} responses.AllProducts
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /cart/add [post]
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

// @Summary Delete from cart
// @Description Delete product form users cart
// @Tags         Cart
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.AddToCart true  "Delete from cart"
// @Success 200 {object} responses.Success
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /cart/delete [delete]
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

		return c.JSON(fiber.Map{
			"result": "success",
		})
	}
}

// @Summary Clear cart
// @Description Delete all product form users cart
// @Tags         Cart
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Clear true  "Delete  all from cart"
// @Success 200 {object} responses.Success
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /cart/clear [delete]
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

		return c.JSON(fiber.Map{
			"result": "success",
		})
	}
}
