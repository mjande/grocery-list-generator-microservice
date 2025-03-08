package handlers

import (
	"encoding/json"
	"log"
	"net/http"

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

	// Temp
	responseData := GroceryListResponse{
		Data: rawIngredients,
	}

	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		log.Println(err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}
