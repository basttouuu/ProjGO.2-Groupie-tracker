package controllers

import (
	"API_HP/src/helpers"
	"API_HP/src/models"
	"API_HP/src/services"
	"API_HP/src/templates"
	"log"
	"net/http"
)

func DisplayListCharacters(w http.ResponseWriter, r *http.Request) {
	data, status, err := services.GetCharactersHomePage()

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des données")
		log.Print(err.Error())
		return
	}

	templates.RenderTemplate(w, r, "home.html", data)
}

func DisplaySearchCharacters(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("search")
	houseFilter := r.URL.Query().Get("house")
	speciesFilter := r.URL.Query().Get("species")
	statusFilter := r.URL.Query().Get("status")
	pageNum := r.URL.Query().Get("page")
	if pageNum == "" {
		pageNum = "1"
	}

	data, status, err := services.GetCharactersSearchPage(searchQuery, houseFilter, speciesFilter, statusFilter, pageNum)

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des données")
		log.Print(err.Error())
		return
	}

	templates.RenderTemplate(w, r, "search.html", data)
}

func DisplayCharacterDetails(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		helpers.RedirectToError(w, r, http.StatusBadRequest, "ID de personnage manquant")
		return
	}

	data, status, err := services.GetCharacterDetailPage(id)

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des données")
		log.Print(err.Error())
		return
	}

	templates.RenderTemplate(w, r, "detail.html", data)
}

func DisplayFavorites(w http.ResponseWriter, r *http.Request) {
	data := services.GetFavoritesPage()
	templates.RenderTemplate(w, r, "favorites.html", data)
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	newFavorite := models.FavoriteCharacter{
		ID:      r.FormValue("id"),
		Name:    r.FormValue("name"),
		Image:   r.FormValue("image"),
		House:   r.FormValue("house"),
		Species: r.FormValue("species"),
	}

	services.AddFavorite(newFavorite)

	returnURL := r.FormValue("return")
	if returnURL == "" {
		returnURL = "/"
	}
	http.Redirect(w, r, returnURL, http.StatusSeeOther)
}

func RemoveFavorite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	services.RemoveFavorite(r.FormValue("id"))

	returnURL := r.FormValue("return")
	if returnURL == "" {
		returnURL = "/"
	}
	http.Redirect(w, r, returnURL, http.StatusSeeOther)
}
