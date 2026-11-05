package models

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Recipe struct {
	ID               int64      `json:"id"`
	Name             string     `json:"name"`
	ShortDescription string     `json:"shortDescription"`
	ImageURL         string     `json:"imageUrl"`
	PrepTimeMinutes  int        `json:"prepTimeMinutes"`
	CookTimeMinutes  int        `json:"cookTimeMinutes"`
	TotalTimeMinutes int        `json:"totalTimeMinutes"`
	Servings         int        `json:"servings"`
	Difficulty       string     `json:"difficulty"`
	Cuisine          string     `json:"cuisine"`
	Categories       []Category `json:"categories"`
}

type RecipeListResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []Recipe `json:"data"`
}
