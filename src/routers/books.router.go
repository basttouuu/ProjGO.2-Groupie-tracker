package routers

import (
	"API_HP/src/controllers"
	"net/http"
)

func BooksRouter(router *http.ServeMux) {
	router.HandleFunc("/books", controllers.DisplayBooks)
	router.HandleFunc("/book/{id}", controllers.DisplayBookDetails)
}
