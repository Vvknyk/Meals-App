package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"meals_backend/config"
	"meals_backend/models"
)

func GetRecipes(w http.ResponseWriter, r *http.Request) {

	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set response content type
	w.Header().Set("Content-Type", "application/json")

	// SQL query
	query := `
        SELECT
            r.id,
            r.name,
            COALESCE(r.short_description, ''),
            COALESCE(r.image_url, ''),
            COALESCE(r.prep_time_minutes, 0),
            COALESCE(r.cook_time_minutes, 0),
            COALESCE(r.servings, 1),
            COALESCE(r.difficulty, ''),
            COALESCE(r.cuisine, ''),
            COALESCE(
                json_agg(
                    json_build_object(
                        'id', c.id,
                        'name', c.name,
                        'slug', c.slug
                    )
                ) FILTER (WHERE c.id IS NOT NULL),
                '[]'::json
            )
        FROM recipes r
        LEFT JOIN recipe_categories rc
            ON r.id = rc.recipe_id
        LEFT JOIN categories c
            ON rc.category_id = c.id
        GROUP BY r.id
        ORDER BY r.id;
    `

	// Execute query
	rows, err := config.DB.Query(r.Context(), query)

	if err != nil {
		log.Println("Database query error:", err)
		http.Error(w, "Failed to fetch recipes", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	recipes := []models.Recipe{}

	// Iterate through database rows
	for rows.Next() {

		var recipe models.Recipe
		var categoriesJSON []byte

		err := rows.Scan(
			&recipe.ID,
			&recipe.Name,
			&recipe.ShortDescription,
			&recipe.ImageURL,
			&recipe.PrepTimeMinutes,
			&recipe.CookTimeMinutes,
			&recipe.Servings,
			&recipe.Difficulty,
			&recipe.Cuisine,
			&categoriesJSON,
		)

		if err != nil {
			log.Println("Row scan error:", err)
			http.Error(w, "Failed to read recipes", http.StatusInternalServerError)
			return
		}

		// Convert category JSON into Go structs
		err = json.Unmarshal(categoriesJSON, &recipe.Categories)

		if err != nil {
			log.Println("Category parsing error:", err)
			http.Error(w, "Failed to process categories", http.StatusInternalServerError)
			return
		}

		recipe.TotalTimeMinutes = recipe.PrepTimeMinutes + recipe.CookTimeMinutes

		recipes = append(recipes, recipe)
	}

	if err := rows.Err(); err != nil {
		log.Println("Rows iteration error:", err)
		http.Error(w, "Failed to fetch recipes", http.StatusInternalServerError)
		return
	}

	// Prepare API response
	response := models.RecipeListResponse{
		Success: true,
		Message: "Recipes fetched successfully",
		Data:    recipes,
	}

	// Return JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
