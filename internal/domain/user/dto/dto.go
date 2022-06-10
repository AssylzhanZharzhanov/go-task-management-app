package dto

import "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user/entity"

type CreateUserDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUserDTO struct {
	ID        entity.UserID `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
}
