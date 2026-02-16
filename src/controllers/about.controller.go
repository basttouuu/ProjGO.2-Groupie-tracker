package controllers

import (
	"API_HP/src/templates"
	"net/http"
)

func AboutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Title": "À Propos - Grimoire des APIs",
		}

		templates.RenderTemplate(w, r, "about.html", data)
	}
}
