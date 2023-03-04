package nosqlstore

import (
	"context"
	"time"

	model "github.com/t67y110v/software-engineering/internal/app/model/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoStoreRepository struct {
	store *Store
}

func (r *MongoStoreRepository) GetProduct() error {

	ctx := context.TODO()

	client := r.store.client

	collection := client.Database("web").Collection("products")

	product := &model.Product{
		ID:                 primitive.NewObjectID(),
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		ProductName:        "api test",
		ProductImgPath:     "test/path",
		ProductPrice:       123,
		ProductDiscount:    12,
		ProductDescription: "description",
	}

	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	return nil

}
