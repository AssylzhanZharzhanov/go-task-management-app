package user

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
)

type Service interface {
	Create(dto user.CreateUserDTO) (int, error)
	List() ([]user.User, error)
	GetByID(userID user.UserID) (user.User, error)
	Update(userDTO user.UpdateUserDTO) (user.User, error)
	Delete(userID user.UserID) error
}
