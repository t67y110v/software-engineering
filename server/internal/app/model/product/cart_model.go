package model

import (
	"testing"
)

type Cart struct {
	UserId string

	ProductName []string `bson:"product_name"`
}

func TestCart(t *testing.T) *Cart {
	return &Cart{
		UserId:      "1",
		ProductName: []string{"test product name", "test product name"},
	}
}
