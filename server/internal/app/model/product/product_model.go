package model

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID                 primitive.ObjectID `bson:"_id"`
	CreatedAt          time.Time          `bson:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at"`
	ProductName        string             `bson:"product_name"`
	ProductCategory    string             `bson:"product_category"`
	ProductImgPath     string             `bson:"product_img_path"`
	ProductPrice       int                `bson:"product_price"`
	ProductDiscount    int                `bson:"product_discount"`
	ProductDescription string             `bson:"product_desccription"`
}

func TestProduct(t *testing.T) *Product {
	return &Product{

		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		ProductName:        "test product name",
		ProductCategory:    "meet",
		ProductImgPath:     "src/img1.jpg",
		ProductPrice:       100,
		ProductDiscount:    0,
		ProductDescription: "very tasty food ",
	}
}

/*
product_name varchar not null,
product_img_path varchar ,
product_price int not null ,
product_discount int DEFAULT 0,
product_description varchar ,
*/
