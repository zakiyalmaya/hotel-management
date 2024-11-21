package model

import "time"

type UserEntity struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	Username  string     `db:"username"`
	Password  string     `db:"password"`
	Email     string     `db:"email"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	Email    string `json:"email" validate:"required,email"`
}

type ChangePasswordRequest struct {
	Username    string `json:"username" validate:"required"`
	OldPassword string `json:"old_password" validate:"required,min=6,max=100"`
	NewPassword string `json:"new_password" validate:"required,min=6,max=100"`
}
