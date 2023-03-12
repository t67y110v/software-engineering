package handlers

import (
	"bytes"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/software-engineering/internal/app/handlers/requests"
)

// @Summary Add product
// @Description Adding new product in the system
// @Tags         Products
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.AddProduct true  "Add product "
// @Success 200 {object} responses.Success
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /product/add [post]
func (h *Handlers) AddProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {

		req := &requests.AddProduct{}

		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle filter by category, status :%d, error :%e", fiber.StatusBadRequest, err)
		}

		if err := h.mgStore.ProductRepository().AddProduct(
			req.ProductName,
			req.ProductCategory,
			req.ProductImgPath,
			req.ProductDescription,
			req.ProductPrice,
			req.ProductDiscount,
		); err != nil {
			return fiber.ErrNotImplemented
		}

		return c.JSON(fiber.Map{
			"result": "success",
		})
	}
}

// @Summary Get all products
// @Description Getting all products
// @Tags         Products
//
//	@Accept       json
//
// @Produce json
// @Success 200 {object} responses.AllProducts
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /product/all [get]
func (h *Handlers) GetAllProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		products, err := h.mgStore.ProductRepository().GetAllProducts()
		if err != nil {
			return err
		}
		return c.JSON(products)
	}
}

// @Summary Filter products by category
// @Description Getting all products in the same categorys
// @Tags         Products
//
//	@Accept       json
//
// @Param        category   path      string  true  "Category"
// @Produce json
// @Success 200 {object} responses.AllProducts
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /product/filter/{category} [get]
func (h *Handlers) FilterByCategory() fiber.Handler {
	return func(c *fiber.Ctx) error {
		category := c.Params("category")

		products, err := h.mgStore.ProductRepository().FilterByCategory(category)
		if err != nil {
			return err
		}

		return c.JSON(products)
	}
}

// @Summary Delete product
// @Description Deleting product by name
// @Tags         Products
//
//	@Accept       json
//
// @Param  data body requests.Delete true  "Delete product "
// @Produce json
// @Success 200 {object} responses.Success
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /product/delete [delete]
func (h *Handlers) DeleteProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {

		req := &requests.Delete{}
		reader := bytes.NewReader(c.Body())
		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle filter by category, status :%d, error :%e", fiber.StatusBadRequest, err)
		}

		if err := h.mgStore.ProductRepository().DeleteProduct(req.Value); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"result": "ok",
		})
	}
}
