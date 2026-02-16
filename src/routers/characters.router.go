package routers

import (
	"API_HP/src/controllers"
	"net/http"
)

func charactersRouter(router *http.ServeMux) {
	router.HandleFunc("/", controllers.DisplayListCharacters)
	router.HandleFunc("/search", controllers.DisplaySearchCharacters)
	router.HandleFunc("/favorites", controllers.DisplayFavorites)
	router.HandleFunc("/favorites/add", controllers.AddFavorite)
	router.HandleFunc("/favorites/remove", controllers.RemoveFavorite)
	router.HandleFunc("/character/{id}", controllers.DisplayCharacterDetails)
}
