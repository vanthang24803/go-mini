package dto

import "time"

type UpdateProfileRequest struct {
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	Phone       string    `json:"phone" validate:"required"`
	Address     string    `json:"address" validate:"required"`
	Gender      string    `json:"gender" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
}
