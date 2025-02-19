package model

import "time"

type AppUser struct {
	UserID           int       `json:"user_id" db:"user_id"`
	Username         string    `json:"username" db:"username"`
	Password         string    `json:"password" db:"password"`
	FullName         string    `json:"full_name" db:"full_name"`
	BirthDate        time.Time `json:"birth_date" db:"birth_date"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
}