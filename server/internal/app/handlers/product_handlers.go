package handlers

import (
	"bytes"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) AddProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			ProductName        string `bson:"product_name"`
			ProductCategory    string `bson:"product_category"`
			ProductImgPath     string `bson:"product_img_path"`
			ProductPrice       int    `bson:"product_price"`
			ProductDiscount    int    `bson:"product_discount"`
			ProductDescription string `bson:"product_desccription"`
		}

		req := &request{}

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

func (h *Handlers) GetAllProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		products, err := h.mgStore.ProductRepository().GetAllProducts()
		if err != nil {
			return err
		}
		return c.JSON(products)
	}
}

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

func (h *Handlers) DeleteProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			Value string `json:"value"`
		}

		req := &request{}
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
