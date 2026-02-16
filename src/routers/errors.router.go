package routers

import (
	"API_HP/src/controllers"
	"net/http"
)

func errorRouter(router *http.ServeMux) {
	router.HandleFunc("/error", controllers.ErrorDisplay)
}
