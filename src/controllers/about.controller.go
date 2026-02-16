package controllers

import (
	"API_HP/src/templates"
	"net/http"
)

func AboutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates.RenderTemplate(w, r, "about.html", nil)
	}
}
