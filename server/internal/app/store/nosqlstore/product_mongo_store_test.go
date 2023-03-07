package nosqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	model "github.com/t67y110v/software-engineering/internal/app/model/product"
	"github.com/t67y110v/software-engineering/internal/app/store/teststore"
)

func TestMongoRepository_AddProduct(t *testing.T) {
	s := teststore.NewMongo()
	p := model.TestProduct(t)

	err := s.ProductRepository().AddProduct(p.ProductName, p.ProductCategory, p.ProductImgPath, p.ProductDescription, p.ProductPrice, p.ProductDiscount)
	assert.NoError(t, err)

	products, err := s.ProductRepository().GetAllProducts()

	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, products[0].ProductName, p.ProductName)

}

func TestMongoRepository_GetAllProducts(t *testing.T) {
	s := teststore.NewMongo()
	p := model.TestProduct(t)

	err := s.ProductRepository().AddProduct(p.ProductName, p.ProductCategory, p.ProductImgPath, p.ProductDescription, p.ProductPrice, p.ProductDiscount)
	assert.NoError(t, err)
	err = s.ProductRepository().AddProduct(p.ProductName, p.ProductCategory, p.ProductImgPath, p.ProductDescription, p.ProductPrice, p.ProductDiscount)
	assert.NoError(t, err)
	err = s.ProductRepository().AddProduct(p.ProductName, p.ProductCategory, p.ProductImgPath, p.ProductDescription, p.ProductPrice, p.ProductDiscount)
	assert.NoError(t, err)

	products, err := s.ProductRepository().GetAllProducts()

	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, len(products), 3)

}

func TestMongoRepository_FilterByCategory(t *testing.T) {
	s := teststore.NewMongo()
	p := model.TestProduct(t)

	testCategory := "test category"
	err := s.ProductRepository().AddProduct(p.ProductName, p.ProductCategory, p.ProductImgPath, p.ProductDescription, p.ProductPrice, p.ProductDiscount)
	assert.NoError(t, err)
	err = s.ProductRepository().AddProduct(p.ProductName, testCategory, p.ProductImgPath, p.ProductDescription, p.ProductPrice, p.ProductDiscount)
	assert.NoError(t, err)
	err = s.ProductRepository().AddProduct(p.ProductName, p.ProductCategory, p.ProductImgPath, p.ProductDescription, p.ProductPrice, p.ProductDiscount)
	assert.NoError(t, err)

	products, err := s.ProductRepository().FilterByCategory(testCategory)
	assert.NoError(t, err)

	assert.Equal(t, products[0].ProductCategory, testCategory)

	products, err = s.ProductRepository().FilterByCategory(p.ProductCategory)
	assert.NoError(t, err)
	assert.Equal(t, len(products), 2)

}

func TestMongoRepository_DeleteProduct(t *testing.T) {
	s := teststore.NewMongo()
	p := model.TestProduct(t)

	err := s.ProductRepository().AddProduct(p.ProductName, p.ProductCategory, p.ProductImgPath, p.ProductDescription, p.ProductPrice, p.ProductDiscount)
	assert.NoError(t, err)

	err = s.ProductRepository().DeleteProduct(p.ProductName)
	assert.NoError(t, err)
	products, _ := s.ProductRepository().GetAllProducts()

	assert.Nil(t, products)

}
