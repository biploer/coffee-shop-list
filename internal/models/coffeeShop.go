package models

import "github.com/google/uuid"

type CoffeeShop struct {
	Uuid        uuid.UUID
	Name        string
	Description string
}

type CreateCoffeeShopInput struct {
	Name        string
	Description string
}
