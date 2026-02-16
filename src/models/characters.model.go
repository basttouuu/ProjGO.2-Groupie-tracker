package models

type APIResponse struct {
	Data []CharacterData `json:"data"`
	Meta Meta            `json:"meta"`
}

type SingleResponse struct {
	Data CharacterData `json:"data"`
}

type CharacterData struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Name       string   `json:"name"`
	House      string   `json:"house"`
	Species    string   `json:"species"`
	Gender     string   `json:"gender"`
	Born       string   `json:"born"`
	Died       string   `json:"died"`
	Patronus   string   `json:"patronus"`
	Image      string   `json:"image"`
	Wiki       string   `json:"wiki"`
	Wands      []string `json:"wands"`
	AliasNames []string `json:"alias_names"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Current int `json:"current"`
	Next    int `json:"next"`
	Last    int `json:"last"`
	Records int `json:"records"`
}

type Character struct {
	ID string
	Attributes
}

func (cd CharacterData) ToCharacter() Character {
	return Character{
		ID:         cd.ID,
		Attributes: cd.Attributes,
	}
}

type FavoriteCharacter struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	House   string `json:"house"`
	Species string `json:"species"`
}
