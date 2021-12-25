package data

//go:generate $GOPATH/bin/easyjson -pkg -no_std_marshalers

import "html/template"

// Pages are ghost pages data
//easyjson:json
type Pages struct {
	Pages []Post `json:"pages"`
	Meta  Meta   `json:"meta"`
}

// Post contains ghost post data
//easyjson:json
type Post struct {
	ID     string        `json:"id"`
	UUID   string        `json:"uuid"`
	Title  string        `json:"title"`
	HTML   template.HTML `json:"html"`
	FImage template.URL  `json:"feature_image"`
}

// Meta contains ghost result metadata
//easyjson:json
type Meta struct {
	Pagination Pagination `json:"pagination"`
}

// Pagination contains ghost pagination data
//easyjson:json
type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Pages int `json:"pages"`
	Total int `json:"total"`
}

// Posts are ghost posts data
//easyjson:json
type Posts struct {
	Posts []Post `json:"posts"`
	Meta  Meta   `json:"meta"`
}
