package routers

import (
	"API_HP/src/controllers"
	"net/http"
)

func AboutRouter(router *http.ServeMux) {
	router.HandleFunc("GET /about", controllers.AboutHandler())
}
