package model

type ContactType struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
type DocumentType struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
