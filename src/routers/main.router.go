package routers

import (
	"API_HP/src/helpers"
	"net/http"
	"path/filepath"
	"runtime"
)

func MainRouter() *http.ServeMux {
	router := http.NewServeMux()

	errorRouter(router)
	charactersRouter(router)
	BooksRouter(router)
	MoviesRouter(router)
	PotionsRouter(router)
	SpellsRouter(router)
	AboutRouter(router)

	_, file, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(file), "..", "..")
	assetsPath := filepath.Join(root, "assets")

	fs := http.FileServer(http.Dir(assetsPath))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	router.HandleFunc("/404", helpers.NotFoundHandler())

	return router
}
