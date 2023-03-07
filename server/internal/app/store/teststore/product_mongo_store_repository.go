package teststore

import (
	"time"

	model "github.com/t67y110v/software-engineering/internal/app/model/product"
	"github.com/t67y110v/software-engineering/internal/app/store"
)

type MongoStoreRepository struct {
	mgStore *mgStore

	products map[int]*model.Product
	carts    map[int]*model.Cart
}

func (r *MongoStoreRepository) AddProduct(name, category, imgPath, description string, price, discount int) error {

	p := &model.Product{

		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		ProductName:        name,
		ProductCategory:    category,
		ProductImgPath:     imgPath,
		ProductDescription: description,
		ProductPrice:       price,
		ProductDiscount:    discount,
	}

	r.products[len(r.products)] = p
	return nil

}

func (r *MongoStoreRepository) GetAllProducts() ([]*model.Product, error) {
	var products []*model.Product
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
}

func (r *MongoStoreRepository) FilterByCategory(category string) ([]*model.Product, error) {

	var products []*model.Product
	for _, product := range r.products {
		if product.ProductCategory == category {
			products = append(products, product)
		}

	}
	return products, nil
}

func (r *MongoStoreRepository) GetProductByName(productName string) (*model.Product, error) {
	for _, product := range r.products {
		if product.ProductName == productName {
			return product, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *MongoStoreRepository) DeleteProduct(productName string) error {

	for id, product := range r.products {
		if product.ProductName == productName {
			delete(r.products, id)
		}
	}
	return nil
}
