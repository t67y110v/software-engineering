package teststore

import (
	model "github.com/t67y110v/software-engineering/internal/app/model/product"
)

func (r *MongoStoreRepository) AddToCart(userId string, productName string) error {
	r.carts[len(r.carts)] = nil
	r.carts[len(r.carts)].UserId = userId
	r.carts[len(r.carts)].ProductName = append(r.carts[len(r.carts)].ProductName, productName)
	return nil

}

func (r *MongoStoreRepository) GetCart(userId string) ([]*model.Product, error) {

	var products []*model.Product

	for _, product := range r.carts {
		if product.UserId == userId {
			for _, prod := range product.ProductName {
				var p model.Product
				p.ProductName = prod
				products = append(products, &p)
			}
		}

	}
	return products, nil
}

func (r *MongoStoreRepository) DeleteFromCart(userId, productName string) error {
	for _, product := range r.carts {
		if product.UserId == userId {
			for _, prod := range product.ProductName {
				if prod == productName {
					prod = ""

				}
			}

		}
	}

	return nil

}

func (r *MongoStoreRepository) ClearCart(userId string) error {
	for _, product := range r.carts {
		if product.UserId == userId {
			product.ProductName = nil
		}
	}
	return nil
}
