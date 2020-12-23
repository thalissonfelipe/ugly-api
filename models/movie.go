package models

// Movie Model
type Movie struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	IMDB     float32   `json:"imdb"`
	Category string    `json:"category"`
	Synosis  string    `json:"synopsis"`
	Director *Director `json:"director"`
}

// Director Model
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Sex       string `json:"sex"`
	Age       int16  `json:"age"`
}
