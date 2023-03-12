package requests

type AddToCart struct {
	UserId  string `json:"user_id"`
	Product string `json:"product_name"`
}
type Clear struct {
	UserId string `json:"user_id"`
}
