//go:generate mockgen -source service.go -destination mock/service.go -package mock
package user

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
)

type Service interface {
	Create(dto *domain.CreateUserDTO) (int, error)
	List() ([]domain.User, error)
	GetByID(userID domain.UserID) (domain.User, error)
	Update(userDTO *domain.UpdateUserDTO) (domain.User, error)
	Delete(userID domain.UserID) error
}
