package routers

import (
	"API_HP/src/controllers"
	"net/http"
)

func SpellsRouter(router *http.ServeMux) {
	router.HandleFunc("/spells", controllers.DisplaySpells)
	router.HandleFunc("/spell/{id}", controllers.DisplaySpellDetails)
}
