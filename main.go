package main

import (
	"net/http"
	"fmt"
	"github.com/Denise3003/RecipeKeeper/data"
	"github.com/Denise3003/RecipeKeeper/handlers"
)

func main() {
	data.FetchAllRecipes()
	fmt.Println(data.AllRecipes)
	http.HandleFunc("/", handlers.HomePage)
	http.ListenAndServe(":8000", nil)
}
