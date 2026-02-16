package services

import (
	"API_HP/src/models"
	"fmt"
	"net/url"
)

func GetAllBooks(query string) ([]models.Book, int, error) {
	var allBooks []models.Book
	currentPage := 1

	for {
		apiURL := fmt.Sprintf("%s/books?page[number]=%d&page[size]=500", PotterDBBaseURL, currentPage)

		if query != "" {
			encodedQuery := url.QueryEscape(query)
			apiURL += fmt.Sprintf("&filter[title_cont]=%s", encodedQuery)
		}

		var response models.BooksAPIResponse
		status, err := callAPI(apiURL, &response)
		if err != nil {
			return allBooks, status, err
		}

		for _, data := range response.Data {
			allBooks = append(allBooks, data.ToBook())
		}

		if response.Meta.Pagination.Next == 0 || response.Meta.Pagination.Current >= response.Meta.Pagination.Last {
			return allBooks, status, nil
		}

		currentPage = response.Meta.Pagination.Next
	}
}

func GetBooks(query, page string) ([]models.Book, *models.Pagination, int, error) {
	apiURL := fmt.Sprintf("%s/books?page[number]=%s&page[size]=500", PotterDBBaseURL, page)

	if query != "" {
		encodedQuery := url.QueryEscape(query)
		apiURL += fmt.Sprintf("&filter[title_cont]=%s", encodedQuery)
	}

	var response models.BooksAPIResponse
	status, err := callAPI(apiURL, &response)
	if err != nil {
		return nil, nil, status, err
	}

	books := make([]models.Book, len(response.Data))
	for i, data := range response.Data {
		books[i] = data.ToBook()
	}

	return books, &response.Meta.Pagination, status, nil
}

func GetBooksPage(query, page string) (models.BooksPageData, int, error) {
	books, pagination, status, err := GetBooks(query, page)
	if err != nil {
		return models.BooksPageData{}, status, err
	}

	return models.BooksPageData{
		PageData: models.PageData{
			Title:      "Livres - Harry Potter Wiki",
			Query:      query,
			Pagination: pagination,
			Page:       page,
		},
		Books: books,
	}, status, nil
}

func GetBookByID(id string) (models.Book, int, error) {
	apiURL := fmt.Sprintf("%s/books/%s", PotterDBBaseURL, id)

	var response models.SingleBookResponse
	status, err := callAPI(apiURL, &response)
	if err != nil {
		return models.Book{}, status, err
	}

	return response.Data.ToBook(), status, nil
}

func GetBookDetail(id string) (models.BookDetailData, int, error) {
	book, status, err := GetBookByID(id)
	if err != nil {
		return models.BookDetailData{}, status, err
	}

	return models.BookDetailData{
		DetailData: models.DetailData{
			Title: "Détails - Harry Potter Wiki",
		},
		Book: book,
	}, status, nil
}

func GetAllSpells(query string) ([]models.Spell, int, error) {
	var allSpells []models.Spell
	currentPage := 1

	for {
		apiURL := fmt.Sprintf("%s/spells?page[number]=%d&page[size]=500", PotterDBBaseURL, currentPage)

		if query != "" {
			encodedQuery := url.QueryEscape(query)
			apiURL += fmt.Sprintf("&filter[name_cont]=%s", encodedQuery)
		}

		var response models.SpellsAPIResponse
		status, err := callAPI(apiURL, &response)
		if err != nil {
			return allSpells, status, err
		}

		for _, data := range response.Data {
			allSpells = append(allSpells, data.ToSpell())
		}

		if response.Meta.Pagination.Next == 0 || response.Meta.Pagination.Current >= response.Meta.Pagination.Last {
			return allSpells, status, nil
		}

		currentPage = response.Meta.Pagination.Next
	}
}

func GetSpells(query, page string) ([]models.Spell, *models.Pagination, int, error) {
	apiURL := fmt.Sprintf("%s/spells?page[number]=%s&page[size]=500", PotterDBBaseURL, page)

	if query != "" {
		encodedQuery := url.QueryEscape(query)
		apiURL += fmt.Sprintf("&filter[name_cont]=%s", encodedQuery)
	}

	var response models.SpellsAPIResponse
	status, err := callAPI(apiURL, &response)
	if err != nil {
		return nil, nil, status, err
	}

	spells := make([]models.Spell, len(response.Data))
	for i, data := range response.Data {
		spells[i] = data.ToSpell()
	}

	return spells, &response.Meta.Pagination, status, nil
}

func GetSpellsPage(query, page string) (models.SpellsPageData, int, error) {
	spells, pagination, status, err := GetSpells(query, page)
	if err != nil {
		return models.SpellsPageData{}, status, err
	}

	return models.SpellsPageData{
		PageData: models.PageData{
			Title:      "Sorts - Harry Potter Wiki",
			Query:      query,
			Pagination: pagination,
			Page:       page,
		},
		Spells: spells,
	}, status, nil
}

func GetSpellByID(id string) (models.Spell, int, error) {
	apiURL := fmt.Sprintf("%s/spells/%s", PotterDBBaseURL, id)

	var response models.SingleSpellResponse
	status, err := callAPI(apiURL, &response)
	if err != nil {
		return models.Spell{}, status, err
	}

	return response.Data.ToSpell(), status, nil
}

func GetSpellDetail(id string) (models.SpellDetailData, int, error) {
	spell, status, err := GetSpellByID(id)
	if err != nil {
		return models.SpellDetailData{}, status, err
	}

	return models.SpellDetailData{
		DetailData: models.DetailData{
			Title: "Détails du sort - Harry Potter Wiki",
		},
		Spell: spell,
	}, status, nil
}

func GetAllPotions(query string) ([]models.Potion, int, error) {
	var allPotions []models.Potion
	currentPage := 1

	for {
		apiURL := fmt.Sprintf("%s/potions?page[number]=%d&page[size]=500", PotterDBBaseURL, currentPage)

		if query != "" {
			encodedQuery := url.QueryEscape(query)
			apiURL += fmt.Sprintf("&filter[name_cont]=%s", encodedQuery)
		}

		var response models.PotionsAPIResponse
		status, err := callAPI(apiURL, &response)
		if err != nil {
			return allPotions, status, err
		}

		for _, data := range response.Data {
			allPotions = append(allPotions, data.ToPotion())
		}

		if response.Meta.Pagination.Next == 0 || response.Meta.Pagination.Current >= response.Meta.Pagination.Last {
			return allPotions, status, nil
		}

		currentPage = response.Meta.Pagination.Next
	}
}

func GetPotions(query, page string) ([]models.Potion, *models.Pagination, int, error) {
	apiURL := fmt.Sprintf("%s/potions?page[number]=%s&page[size]=500", PotterDBBaseURL, page)

	if query != "" {
		encodedQuery := url.QueryEscape(query)
		apiURL += fmt.Sprintf("&filter[name_cont]=%s", encodedQuery)
	}

	var response models.PotionsAPIResponse
	status, err := callAPI(apiURL, &response)
	if err != nil {
		return nil, nil, status, err
	}

	potions := make([]models.Potion, len(response.Data))
	for i, data := range response.Data {
		potions[i] = data.ToPotion()
	}

	return potions, &response.Meta.Pagination, status, nil
}

func GetPotionsPage(query, page string) (models.PotionsPageData, int, error) {
	potions, pagination, status, err := GetPotions(query, page)
	if err != nil {
		return models.PotionsPageData{}, status, err
	}

	return models.PotionsPageData{
		PageData: models.PageData{
			Title:      "Potions - Harry Potter Wiki",
			Query:      query,
			Pagination: pagination,
			Page:       page,
		},
		Potions: potions,
	}, status, nil
}

func GetPotionByID(id string) (models.Potion, int, error) {
	apiURL := fmt.Sprintf("%s/potions/%s", PotterDBBaseURL, id)

	var response models.SinglePotionResponse
	status, err := callAPI(apiURL, &response)
	if err != nil {
		return models.Potion{}, status, err
	}

	return response.Data.ToPotion(), status, nil
}

func GetPotionDetail(id string) (models.PotionDetailData, int, error) {
	potion, status, err := GetPotionByID(id)
	if err != nil {
		return models.PotionDetailData{}, status, err
	}

	return models.PotionDetailData{
		DetailData: models.DetailData{
			Title: "Détails de la potion - Harry Potter Wiki",
		},
		Potion: potion,
	}, status, nil
}

func GetAllMovies(query string) ([]models.Movie, int, error) {
	var allMovies []models.Movie
	currentPage := 1

	for {
		apiURL := fmt.Sprintf("%s/movies?page[number]=%d&page[size]=500", PotterDBBaseURL, currentPage)

		if query != "" {
			encodedQuery := url.QueryEscape(query)
			apiURL += fmt.Sprintf("&filter[title_cont]=%s", encodedQuery)
		}

		var response models.MoviesAPIResponse
		status, err := callAPI(apiURL, &response)
		if err != nil {
			return allMovies, status, err
		}

		for _, data := range response.Data {
			allMovies = append(allMovies, data.ToMovie())
		}

		if response.Meta.Pagination.Next == 0 || response.Meta.Pagination.Current >= response.Meta.Pagination.Last {
			return allMovies, status, nil
		}

		currentPage = response.Meta.Pagination.Next
	}
}

func GetMovies(query, page string) ([]models.Movie, *models.Pagination, int, error) {
	apiURL := fmt.Sprintf("%s/movies?page[number]=%s&page[size]=500", PotterDBBaseURL, page)

	if query != "" {
		encodedQuery := url.QueryEscape(query)
		apiURL += fmt.Sprintf("&filter[title_cont]=%s", encodedQuery)
	}

	var response models.MoviesAPIResponse
	status, err := callAPI(apiURL, &response)
	if err != nil {
		return nil, nil, status, err
	}

	movies := make([]models.Movie, len(response.Data))
	for i, data := range response.Data {
		movies[i] = data.ToMovie()
	}

	return movies, &response.Meta.Pagination, status, nil
}

func GetMoviesPage(query, page string) (models.MoviesPageData, int, error) {
	movies, pagination, status, err := GetMovies(query, page)
	if err != nil {
		return models.MoviesPageData{}, status, err
	}

	return models.MoviesPageData{
		PageData: models.PageData{
			Title:      "Films - Harry Potter Wiki",
			Query:      query,
			Pagination: pagination,
			Page:       page,
		},
		Movies: movies,
	}, status, nil
}

func GetMovieByID(id string) (models.Movie, int, error) {
	apiURL := fmt.Sprintf("%s/movies/%s", PotterDBBaseURL, id)

	var response models.SingleMovieResponse
	status, err := callAPI(apiURL, &response)
	if err != nil {
		return models.Movie{}, status, err
	}

	return response.Data.ToMovie(), status, nil
}

func GetMovieDetail(id string) (models.MovieDetailData, int, error) {
	movie, status, err := GetMovieByID(id)
	if err != nil {
		return models.MovieDetailData{}, status, err
	}

	return models.MovieDetailData{
		DetailData: models.DetailData{
			Title: "Détails du film - Harry Potter Wiki",
		},
		Movie: movie,
	}, status, nil
}
