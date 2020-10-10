package models

// Book structure
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author structure
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
