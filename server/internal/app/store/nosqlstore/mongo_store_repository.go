package nosqlstore

import (
	"context"
	"time"

	model "github.com/t67y110v/software-engineering/internal/app/model/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStoreRepository struct {
	store *Store
}

func (r *MongoStoreRepository) GetProduct() error {

	ctx := context.TODO()

	collection := r.store.client.Database("web").Collection("products")

	product := &model.Product{
		ID:                 primitive.NewObjectID(),
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		ProductName:        "api test2",
		ProductCategory:    "tasty",
		ProductImgPath:     "test/path//1",
		ProductPrice:       123,
		ProductDiscount:    12,
		ProductDescription: "description very interesting ",
	}

	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	return nil

}

func (r *MongoStoreRepository) GetAllProducts() ([]*model.Product, error) {
	filter := bson.D{{}}
	ctx := context.TODO()
	collection := r.store.client.Database("web").Collection("products")

	cur, err := collection.Find(ctx, filter)

	var products []*model.Product

	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var p model.Product
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	if err := cur.Err(); err != nil {
		return products, err
	}
	cur.Close(ctx)
	if len(products) == 0 {
		return products, mongo.ErrNoDocuments

	}
	return products, nil

}

func (r MongoStoreRepository) Filter(category string) ([]*model.Product, error) {
	filter := bson.D{
		primitive.E{
			Key:   "product_category",
			Value: category,
		},
	}
	ctx := context.TODO()
	collection := r.store.client.Database("web").Collection("products")

	cur, err := collection.Find(ctx, filter)

	var products []*model.Product

	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var p model.Product
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	if err := cur.Err(); err != nil {
		return products, err
	}
	cur.Close(ctx)
	if len(products) == 0 {
		return products, mongo.ErrNoDocuments

	}
	return products, nil

}
