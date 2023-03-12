package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Success struct {
	Result string `json:"result"`
}

type AllProducts struct {
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
