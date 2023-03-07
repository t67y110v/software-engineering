package nosqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	model "github.com/t67y110v/software-engineering/internal/app/model/product"
	"github.com/t67y110v/software-engineering/internal/app/store/teststore"
)

func TestMongoRepository_AddToCart(t *testing.T) {
	s := teststore.NewMongo()
	c := model.TestCart(t)

	err := s.ProductRepository().AddToCart(c.UserId, c.ProductName[0])
	assert.NoError(t, err)

}
