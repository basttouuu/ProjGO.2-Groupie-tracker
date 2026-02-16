package routers

import (
	"API_HP/src/controllers"
	"net/http"
)

func MoviesRouter(router *http.ServeMux) {
	router.HandleFunc("/movies", controllers.DisplayMovies)
	router.HandleFunc("/movie/{id}", controllers.DisplayMovieDetails)
}
