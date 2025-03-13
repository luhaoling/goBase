package dataStruct

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

type CartItem struct {
	ProductID   int
	ProductName string
	Quantity    int
	Price       float64
}

type Product struct {
	ID          int
	ProductName string
	Price       float64
}

func getProductByID(productId int) Product {
	return Product{
		ID:          1,
		ProductName: "product",
		Price:       190,
	}
}

func AddToCart(userID, productId int) {
	product := getProductByID(productId)

	item := &CartItem{
		ProductID:   product.ID,
		ProductName: product.ProductName,
		Quantity:    0,
		Price:       product.Price,
	}

	data, err := json.Marshal(item)
	if err != nil {
		log.Println("failed to marshal item:", err)
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	err = client.RPush(context.Background(), fmt.Sprintf("cart:%d", userID), data).Err()
	if err != nil {
		log.Println("failed to add cart item:", err)
		return
	}

	log.Println("cart item add successfully")
}
