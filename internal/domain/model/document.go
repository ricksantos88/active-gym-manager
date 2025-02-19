package model

import "time"

type Document struct {
	ID             int       `json:"id" db:"id"`
	UserID         int       `json:"user_id" db:"user_id"`
	DocumentTypeID int       `json:"document_type_id" db:"document_type_id"`
	Value          string    `json:"value" db:"value"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	CreatedBy      int       `json:"created_by" db:"created_by"`
	UpdatedBy      int       `json:"updated_by" db:"updated_by"`
}