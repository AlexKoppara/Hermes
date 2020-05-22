// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type MenuItem struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PriceInCents int    `json:"priceInCents"`
}

func (MenuItem) IsEntity() {}

type Order struct {
	ID          string       `json:"id"`
	MenuItems   []*MenuItem  `json:"menuItems"`
	Total       int          `json:"total"`
	OrderStatus *OrderStatus `json:"orderStatus"`
}

type OrderStatus struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
