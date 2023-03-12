package requests

type AddProduct struct {
	ProductName        string `bson:"product_name"`
	ProductCategory    string `bson:"product_category"`
	ProductImgPath     string `bson:"product_img_path"`
	ProductPrice       int    `bson:"product_price"`
	ProductDiscount    int    `bson:"product_discount"`
	ProductDescription string `bson:"product_desccription"`
}

type Delete struct {
	Value string `json:"value"`
}
