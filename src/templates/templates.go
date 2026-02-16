package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

var listTemp *template.Template

func Load() {
	listTemp = template.Must(template.New("").Funcs(template.FuncMap{
		"add":      func(a, b int) int { return a + b },
		"sub":      func(a, b int) int { return a - b },
		"safeHTML": func(s string) template.HTML { return template.HTML(s) },
	}).ParseGlob("../../templates/*.html"))
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	var buffer bytes.Buffer
	err := listTemp.ExecuteTemplate(&buffer, name, data)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=%s", http.StatusInternalServerError, url.QueryEscape("Error loading page")), http.StatusSeeOther)
		return
	}
	buffer.WriteTo(w)
}
