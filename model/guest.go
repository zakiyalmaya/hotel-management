package model

import "time"

type GuestEntity struct {
	ID          int        `db:"id"`
	FirstName   string     `db:"first_name"`
	LastName    string     `db:"last_name"`
	Identity    string     `db:"identity_number"`
	DateOfBirth time.Time  `db:"date_of_birth"`
	PhoneNumber string     `db:"phone_number"`
	Email       string     `db:"email"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type GuestResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Identity    string `json:"identity_number"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type CreateGuestRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Identity    string `json:"identity_number" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"required"`
}
