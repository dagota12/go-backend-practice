package models

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	// Status indicates whether the book is available or not. If true, the book
	// is available; otherwise, it is unavailable.
	Status bool `json:"status"`
}
