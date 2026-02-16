package controllers

import (
	"API_HP/src/helpers"
	"API_HP/src/services"
	"API_HP/src/templates"
	"log"
	"net/http"
)

func DisplayBooks(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("search")
	pageNum := r.URL.Query().Get("page")
	if pageNum == "" {
		pageNum = "1"
	}

	data, status, err := services.GetBooksPage(searchQuery, pageNum)

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des livres")
		log.Print(err.Error())
		return
	}

	templates.RenderTemplate(w, r, "books.html", data)
}

func DisplayBookDetails(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		helpers.RedirectToError(w, r, http.StatusBadRequest, "ID de livre manquant")
		return
	}

	data, status, err := services.GetBookDetail(id)

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération du livre")
		log.Print(err.Error())
		return
	}

	templates.RenderTemplate(w, r, "book-detail.html", data)
}

func DisplaySpells(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("search")
	pageNum := r.URL.Query().Get("page")
	if pageNum == "" {
		pageNum = "1"
	}

	data, status, err := services.GetSpellsPage(searchQuery, pageNum)

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des sorts")
		log.Print(err.Error())
		return
	}

	templates.RenderTemplate(w, r, "spells.html", data)
}

func DisplaySpellDetails(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		helpers.RedirectToError(w, r, http.StatusBadRequest, "ID de sort manquant")
		return
	}

	data, status, err := services.GetSpellDetail(id)

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération du sort")
		log.Print(err.Error())
		return
	}

	templates.RenderTemplate(w, r, "spell-detail.html", data)
}

func DisplayMovies(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("search")
	pageNum := r.URL.Query().Get("page")
	if pageNum == "" {
		pageNum = "1"
	}

	data, status, err := services.GetMoviesPage(searchQuery, pageNum)

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des films")
		log.Print(err.Error())
		return
	}

	templates.RenderTemplate(w, r, "movies.html", data)
}

func DisplayMovieDetails(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		helpers.RedirectToError(w, r, http.StatusBadRequest, "ID de film manquant")
		return
	}

	data, status, err := services.GetMovieDetail(id)

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération du film")
		log.Print(err.Error())
		return
	}

	templates.RenderTemplate(w, r, "movie-detail.html", data)
}

func DisplayPotions(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("search")
	pageNum := r.URL.Query().Get("page")
	if pageNum == "" {
		pageNum = "1"
	}

	data, status, err := services.GetPotionsPage(searchQuery, pageNum)

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des potions")
		log.Print(err.Error())
		return
	}

	templates.RenderTemplate(w, r, "potions.html", data)
}

func DisplayPotionDetails(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		helpers.RedirectToError(w, r, http.StatusBadRequest, "ID de potion manquant")
		return
	}

	data, status, err := services.GetPotionDetail(id)

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération de la potion")
		log.Print(err.Error())
		return
	}

	templates.RenderTemplate(w, r, "potion-detail.html", data)
}
