package controllers

import (
	"API_HP/src/models"
	"API_HP/src/templates"
	"net/http"
	"strconv"
)

func ErrorDisplay(w http.ResponseWriter, r *http.Request) {
	codeStr := r.FormValue("code")
	statusCode := 500

	if codeStr != "" {
		if code, err := strconv.Atoi(codeStr); err == nil {
			statusCode = code
		}
	}

	data := models.Error{
		StatusCode:   statusCode,
		Message:      r.FormValue("message"),
		ErrorDetails: r.FormValue("details"),
	}

	w.WriteHeader(statusCode)
	templates.RenderTemplate(w, r, "error.html", data)
}
