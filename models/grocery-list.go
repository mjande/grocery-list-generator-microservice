package models

type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Units    string  `json:"units"`
}

type GroceryList []Ingredient
