package user

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
)

type PostgresRepository interface {
	Create(user user.User) (int, error)
	List() ([]user.User, error)
	GetByID(userID user.UserID) (user.User, error)
	Update(user user.User) (user.User, error)
	Delete(userID user.UserID) error
}
