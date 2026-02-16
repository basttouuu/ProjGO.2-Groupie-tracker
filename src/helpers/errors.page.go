package helpers

import (
	"API_HP/src/templates"
	"net/http"
)

func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		data := struct {
			StatusCode int
			Message    string
		}{
			StatusCode: http.StatusNotFound,
			Message:    "",
		}
		templates.RenderTemplate(w, r, "error.html", data)
	}
}
