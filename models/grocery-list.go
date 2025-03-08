package models

import "strings"

// Represents ingredient with standard fields.
type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

// Represents a variant of Ingredient which has all different
// units (with quantities) combined into a slice.
type GroceryIngredient struct {
	Name       string       `json:"name"`
	Quantities []GroceryQty `json:"quantities"`
}

// Represents a unit/quantity pair. Used for representing different
// units for the same ingredient.
type GroceryQty struct {
	Unit     string  `json:"unit"`
	Quantity float64 `json:"quantity"`
}

// Represents a collection of ingredients that comprise a grocery list.
type GroceryList []GroceryIngredient

func CondenseIntoGroceryList(ingredients []Ingredient) GroceryList {
	ingredientsMap := map[string][]Ingredient{}

	// Put ingredients with similar names together in map
	for _, ingredient := range ingredients {
		ingredientName := strings.ToLower(ingredient.Name)

		prevIngredients, exists := ingredientsMap[ingredientName]
		if exists {
			ingredientsMap[ingredientName] = append(prevIngredients, ingredient)
		} else {
			ingredientsMap[ingredientName] = []Ingredient{ingredient}
		}
	}

	// Create grocery list
	groceryList := GroceryList{}
	for _, similarIngredients := range ingredientsMap {
		combinedIngredient := combineSimilarIngredients(similarIngredients)

		groceryList = append(groceryList, combinedIngredient)
	}

	return groceryList
}

// Helper functions

// Takes a list of ingredients with the same name but possibly
// different units. Combines similar units using addition, and collects
// different units into a slice.
func combineSimilarIngredients(ingredients []Ingredient) GroceryIngredient {
	ingredientName := ingredients[0].Name

	// Map similar units
	unitMap := map[string]float64{}
	for _, ingredient := range ingredients {
		unit := strings.ToLower(ingredient.Unit)

		_, exists := unitMap[unit]
		if exists {
			unitMap[unit] += ingredient.Quantity
		} else {
			unitMap[unit] = ingredient.Quantity
		}
	}

	groceryIngredient := GroceryIngredient{
		Name:       ingredientName,
		Quantities: []GroceryQty{},
	}

	// Collects different units into slice
	for unit, qty := range unitMap {
		groceryIngredient.Quantities = append(groceryIngredient.Quantities, GroceryQty{unit, qty})
	}

	return groceryIngredient
}
