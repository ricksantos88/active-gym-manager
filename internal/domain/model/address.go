package model

import "time"

type Address struct {
	ID           int       `json:"id" db:"id"`
	UserID       int       `json:"user_id" db:"user_id"`
	Street       string    `json:"street" db:"street"`
	Number       string    `json:"number" db:"number"`
	Complement   string    `json:"complement" db:"complement"`
	Neighborhood string    `json:"neighborhood" db:"neighborhood"`
	City         string    `json:"city" db:"city"`
	State        string    `json:"state" db:"state"`
	ZipCode      string    `json:"zip_code" db:"zip_code"`
	Country      string    `json:"country" db:"country"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	CreatedBy    int       `json:"created_by" db:"created_by"`
	UpdatedBy    int       `json:"updated_by" db:"updated_by"`
}