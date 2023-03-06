package nosqlstore

import (
	"context"
	"errors"

	model "github.com/t67y110v/software-engineering/internal/app/model/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *MongoStoreRepository) AddToCart(userId string, productName string) error {
	ctx := context.TODO()

	collection := r.store.client.Database("web").Collection("cart")

	newItem := bson.D{
		{
			Key:   "user_id",
			Value: userId,
		},
		{
			Key:   "product_name",
			Value: productName,
		},
	}

	_, err := collection.InsertOne(ctx, newItem)
	if err != nil {
		return err
	}

	return nil

}

func (r *MongoStoreRepository) GetCart(userId string) ([]*model.Product, error) {

	filter := bson.D{
		primitive.E{
			Key:   "user_id",
			Value: userId,
		},
	}

	ctx := context.TODO()
	collection := r.store.client.Database("web").Collection("cart")

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
		pr, err := r.store.mongoStoreRepository.GetProductByName(p.ProductName)
		if err != nil {
			return nil, err
		}
		products = append(products, pr)
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

func (r *MongoStoreRepository) DeleteFromCart(userId, productName string) error {

	ctx := context.TODO()
	filter := bson.D{
		primitive.E{
			Key:   "user_id",
			Value: userId,
		},
		{
			Key:   "product_name",
			Value: productName,
		},
	}

	collection := r.store.client.Database("web").Collection("cart")

	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no products were deleted")
	}
	return nil

}

func (r *MongoStoreRepository) ClearCart(userId string) error {
	ctx := context.TODO()
	filter := bson.D{primitive.E{
		Key:   "user_id",
		Value: userId,
	}}

	collection := r.store.client.Database("web").Collection("cart")

	res, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no products were deleted")
	}
	return nil

}
