package routers

import (
	"API_HP/src/controllers"
	"net/http"
)

func PotionsRouter(router *http.ServeMux) {
	router.HandleFunc("/potions", controllers.DisplayPotions)
	router.HandleFunc("/potion/{id}", controllers.DisplayPotionDetails)
}
