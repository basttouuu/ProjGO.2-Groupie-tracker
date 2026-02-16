package models

type BooksAPIResponse struct {
	Data []BookData `json:"data"`
	Meta Meta       `json:"meta"`
}

type SingleBookResponse struct {
	Data BookData `json:"data"`
}

type BookData struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes BookAttributes `json:"attributes"`
}

type BookAttributes struct {
	Slug        string `json:"slug"`
	Author      string `json:"author"`
	Cover       string `json:"cover"`
	Dedication  string `json:"dedication"`
	Pages       int    `json:"pages"`
	ReleaseDate string `json:"release_date"`
	Summary     string `json:"summary"`
	Title       string `json:"title"`
	Wiki        string `json:"wiki"`
}

type Book struct {
	ID string
	BookAttributes
}

func (bd BookData) ToBook() Book {
	book := Book{
		ID:             bd.ID,
		BookAttributes: bd.Attributes,
	}
	if book.Cover == "" {
		book.Cover = "/static/img/no-image-2.png"
	}
	return book
}

type SpellsAPIResponse struct {
	Data []SpellData `json:"data"`
	Meta Meta        `json:"meta"`
}

type SingleSpellResponse struct {
	Data SpellData `json:"data"`
}

type SpellData struct {
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	Attributes SpellAttributes `json:"attributes"`
}

type SpellAttributes struct {
	Slug        string `json:"slug"`
	Category    string `json:"category"`
	Creator     string `json:"creator"`
	Effect      string `json:"effect"`
	Hand        string `json:"hand"`
	Image       string `json:"image"`
	Incantation string `json:"incantation"`
	Light       string `json:"light"`
	Name        string `json:"name"`
	Wiki        string `json:"wiki"`
}

type Spell struct {
	ID string
	SpellAttributes
}

func (sd SpellData) ToSpell() Spell {
	spell := Spell{
		ID:              sd.ID,
		SpellAttributes: sd.Attributes,
	}
	if spell.Image == "" {
		spell.Image = "/static/img/no-image-2.png"
	}
	return spell
}

type MoviesAPIResponse struct {
	Data []MovieData `json:"data"`
	Meta Meta        `json:"meta"`
}

type SingleMovieResponse struct {
	Data MovieData `json:"data"`
}

type MovieData struct {
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	Attributes MovieAttributes `json:"attributes"`
}

type MovieAttributes struct {
	Slug             string   `json:"slug"`
	BoxOffice        string   `json:"box_office"`
	Budget           string   `json:"budget"`
	Cinematographers []string `json:"cinematographers"`
	Directors        []string `json:"directors"`
	Distributors     []string `json:"distributors"`
	Editors          []string `json:"editors"`
	MusicComposers   []string `json:"music_composers"`
	Poster           string   `json:"poster"`
	Producers        []string `json:"producers"`
	ReleaseDate      string   `json:"release_date"`
	RunningTime      string   `json:"running_time"`
	Screenwriters    []string `json:"screenwriters"`
	Summary          string   `json:"summary"`
	Title            string   `json:"title"`
	Trailer          string   `json:"trailer"`
	Wiki             string   `json:"wiki"`
}

type Movie struct {
	ID string
	MovieAttributes
}

func (md MovieData) ToMovie() Movie {
	movie := Movie{
		ID:              md.ID,
		MovieAttributes: md.Attributes,
	}
	if movie.Poster == "" {
		movie.Poster = "/static/img/no-image-2.png"
	}
	return movie
}

type PotionsAPIResponse struct {
	Data []PotionData `json:"data"`
	Meta Meta         `json:"meta"`
}

type SinglePotionResponse struct {
	Data PotionData `json:"data"`
}

type PotionData struct {
	ID         string           `json:"id"`
	Type       string           `json:"type"`
	Attributes PotionAttributes `json:"attributes"`
}

type PotionAttributes struct {
	Slug            string `json:"slug"`
	Characteristics string `json:"characteristics"`
	Difficulty      string `json:"difficulty"`
	Effect          string `json:"effect"`
	Image           string `json:"image"`
	Ingredients     string `json:"ingredients"`
	Inventors       string `json:"inventors"`
	Manufacturers   string `json:"manufacturers"`
	Name            string `json:"name"`
	SideEffects     string `json:"side_effects"`
	Time            string `json:"time"`
	Wiki            string `json:"wiki"`
}

type Potion struct {
	ID string
	PotionAttributes
}

func (pd PotionData) ToPotion() Potion {
	potion := Potion{
		ID:               pd.ID,
		PotionAttributes: pd.Attributes,
	}
	if potion.Image == "" {
		potion.Image = "/static/img/no-image-2.png"
	}
	return potion
}
