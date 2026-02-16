package services

import (
	"API_HP/src/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
)

var (
	favoritesMux sync.Mutex
)

func getFavoritesFilePath() string {
	return "favorites.json"
}

func GetFavorites() []models.FavoriteCharacter {
	favoritesMux.Lock()
	defer favoritesMux.Unlock()

	data, err := os.ReadFile(getFavoritesFilePath())
	if err != nil {
		return []models.FavoriteCharacter{}
	}

	var favorites []models.FavoriteCharacter
	json.Unmarshal(data, &favorites)
	return favorites
}

func SaveFavorites(favorites []models.FavoriteCharacter) error {
	favoritesMux.Lock()
	defer favoritesMux.Unlock()

	data, err := json.MarshalIndent(favorites, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(getFavoritesFilePath(), data, 0644)
}

func IsFavorite(favorites []models.FavoriteCharacter, id string) bool {
	for _, fav := range favorites {
		if fav.ID == id {
			return true
		}
	}
	return false
}

func buildCharactersURL(query, house, species, statusFilter, page string) string {
	apiURL := fmt.Sprintf("%s/characters?page[number]=%s&page[size]=500", PotterDBBaseURL, page)

	if query != "" {
		encodedQuery := url.QueryEscape(query)
		apiURL += fmt.Sprintf("&filter[name_cont]=%s", encodedQuery)
	}

	if house != "" {
		encodedHouse := url.QueryEscape(house)
		apiURL += fmt.Sprintf("&filter[house_eq]=%s", encodedHouse)
	}

	if species != "" {
		encodedSpecies := url.QueryEscape(species)
		apiURL += fmt.Sprintf("&filter[species_eq]=%s", encodedSpecies)
	}

	if statusFilter != "" {
		if statusFilter == "alive" {
			apiURL += "&filter[died_null]=1"
		} else if statusFilter == "deceased" {
			apiURL += "&filter[died_not_null]=1"
		}
	}

	return apiURL
}

func GetAllCharacters(query, house, species, statusFilter string) ([]models.Character, int, error) {
	var allCharacters []models.Character
	currentPage := 1

	for {
		apiURL := buildCharactersURL(query, house, species, statusFilter, fmt.Sprintf("%d", currentPage))

		var response models.APIResponse
		status, err := callAPI(apiURL, &response)
		if err != nil {
			return allCharacters, status, err
		}

		for _, data := range response.Data {
			allCharacters = append(allCharacters, data.ToCharacter())
		}

		if response.Meta.Pagination.Next == 0 || response.Meta.Pagination.Current >= response.Meta.Pagination.Last {
			return allCharacters, status, nil
		}

		currentPage = response.Meta.Pagination.Next
	}
}

func GetCharacters(query, house, species, statusFilter, page string) ([]models.Character, *models.Pagination, int, error) {
	apiURL := buildCharactersURL(query, house, species, statusFilter, page)

	var response models.APIResponse
	status, err := callAPI(apiURL, &response)
	if err != nil {
		return nil, nil, status, err
	}

	characters := make([]models.Character, len(response.Data))
	for i, data := range response.Data {
		characters[i] = data.ToCharacter()
	}

	return characters, &response.Meta.Pagination, status, nil
}

func GetCharacterByID(id string) (models.Character, int, error) {
	apiURL := fmt.Sprintf("%s/characters/%s", PotterDBBaseURL, id)

	var response models.SingleResponse
	status, err := callAPI(apiURL, &response)
	if err != nil {
		return models.Character{}, status, err
	}

	return response.Data.ToCharacter(), status, nil
}

func GetCharactersHomePage() (models.CharactersHomePageData, int, error) {
	mainCharacterIDs := []string{
		"harry-potter", "hermione-granger", "ron-weasley",
		"albus-dumbledore", "severus-snape", "voldemort",
	}

	var mainCharacters []models.Character
	for _, id := range mainCharacterIDs {
		character, status, err := GetCharacterByID(id)
		if status == http.StatusOK && err == nil {
			mainCharacters = append(mainCharacters, character)
		}
	}

	return models.CharactersHomePageData{
		Title:      "Harry Potter Wiki - Grimoire Céleste",
		Characters: mainCharacters,
		Favorites:  GetFavorites(),
	}, http.StatusOK, nil
}

func GetCharactersSearchPage(query, house, species, statusFilter, page string) (models.CharactersSearchPageData, int, error) {
	characters, pagination, status, err := GetCharacters(query, house, species, statusFilter, page)
	if err != nil {
		return models.CharactersSearchPageData{}, status, err
	}

	return models.CharactersSearchPageData{
		Title:      "Recherche - Harry Potter Wiki",
		Query:      query,
		House:      house,
		Species:    species,
		Status:     statusFilter,
		Characters: characters,
		Pagination: pagination,
		Page:       page,
		Favorites:  GetFavorites(),
	}, status, nil
}

func GetCharacterDetailPage(id string) (models.CharacterDetailData, int, error) {
	character, status, err := GetCharacterByID(id)
	if err != nil {
		return models.CharacterDetailData{}, status, err
	}

	return models.CharacterDetailData{
		Title:     "Détails - Harry Potter Wiki",
		Character: character,
		Favorites: GetFavorites(),
	}, status, nil
}

func GetFavoritesPage() models.FavoritesPageData {
	return models.FavoritesPageData{
		Title:      "Harry Potter Wiki - Favoris",
		Characters: GetFavorites(),
	}
}

func AddFavorite(favorite models.FavoriteCharacter) error {
	favorites := GetFavorites()

	if IsFavorite(favorites, favorite.ID) {
		return nil
	}

	favorites = append(favorites, favorite)
	return SaveFavorites(favorites)
}

func RemoveFavorite(id string) error {
	favorites := GetFavorites()
	newFavorites := []models.FavoriteCharacter{}

	for _, fav := range favorites {
		if fav.ID != id {
			newFavorites = append(newFavorites, fav)
		}
	}

	return SaveFavorites(newFavorites)
}
