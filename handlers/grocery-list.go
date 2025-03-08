package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"

	"github.com/mjande/grocery-list-generator-microservice/models"
	"github.com/mjande/grocery-list-generator-microservice/utils"
)

type GroceryListResponse struct {
	Message string             `json:"message"`
	Data    models.GroceryList `json:"data"`
}

func PostGroceryList(w http.ResponseWriter, r *http.Request) {
	var rawIngredients []models.Ingredient
	err := json.NewDecoder(r.Body).Decode(&rawIngredients)
	if err != nil {
		log.Println(err)
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	groceryList := models.CondenseIntoGroceryList(rawIngredients)

	sort.Slice(groceryList, func(i, j int) bool {
		return groceryList[i].Name < groceryList[j].Name
	})

	responseData := GroceryListResponse{
		Data: groceryList,
	}

	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		log.Println(err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}
