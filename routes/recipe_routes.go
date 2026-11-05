package routes

import (
	"net/http"

	"meals_backend/handlers"
)

func RegisterRecipeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/recipes", handlers.GetRecipes)
}
