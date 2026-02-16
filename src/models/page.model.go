package models

type PageData struct {
	Title      string
	Query      string
	Pagination *Pagination
	Page       string
}

type DetailData struct {
	Title string
}

type BooksPageData struct {
	PageData
	Books []Book
}

type BookDetailData struct {
	DetailData
	Book Book
}

type SpellsPageData struct {
	PageData
	Spells []Spell
}

type SpellDetailData struct {
	DetailData
	Spell Spell
}

type PotionsPageData struct {
	PageData
	Potions []Potion
}

type PotionDetailData struct {
	DetailData
	Potion Potion
}

type MoviesPageData struct {
	PageData
	Movies []Movie
}

type MovieDetailData struct {
	DetailData
	Movie Movie
}

type CharactersHomePageData struct {
	Title      string
	Characters []Character
	Favorites  []FavoriteCharacter
}

type CharactersSearchPageData struct {
	Title      string
	Query      string
	House      string
	Species    string
	Status     string
	Characters []Character
	Pagination *Pagination
	Page       string
	Favorites  []FavoriteCharacter
}

type CharacterDetailData struct {
	Title     string
	Character Character
	Favorites []FavoriteCharacter
}

type FavoritesPageData struct {
	Title      string
	Characters []FavoriteCharacter
}
