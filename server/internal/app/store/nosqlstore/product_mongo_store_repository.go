package nosqlstore

import (
	"context"
	"errors"
	"time"

	model "github.com/t67y110v/software-engineering/internal/app/model/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStoreRepository struct {
	store *Store
}

func (r *MongoStoreRepository) AddProduct(name, category, imgPath, description string, price, discount int) error {

	ctx := context.TODO()

	collection := r.store.client.Database("web").Collection("products")

	product := &model.Product{
		ID:                 primitive.NewObjectID(),
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		ProductName:        name,
		ProductCategory:    category,
		ProductImgPath:     imgPath,
		ProductPrice:       price,
		ProductDiscount:    discount,
		ProductDescription: description,
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

func (r *MongoStoreRepository) FilterByCategory(category string) ([]*model.Product, error) {
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

func (r *MongoStoreRepository) GetProductByName(productName string) (*model.Product, error) {
	filter := bson.D{
		primitive.E{
			Key:   "product_name",
			Value: productName,
		},
	}
	ctx := context.TODO()
	collection := r.store.client.Database("web").Collection("products")

	cur := collection.FindOne(ctx, filter)

	var p model.Product

	err := cur.Decode(&p)
	if err := cur.Err(); err != nil {
		return &p, err
	}
	if err != nil {
		return nil, err
	}

	return &p, nil

}

func (r *MongoStoreRepository) DeleteProduct(productName string) error {
	ctx := context.TODO()
	filter := bson.D{primitive.E{
		Key:   "product_name",
		Value: productName,
	}}

	collection := r.store.client.Database("web").Collection("products")

	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no products were deleted")
	}
	return nil

}
