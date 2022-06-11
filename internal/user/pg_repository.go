//go:generate mockgen -source pg_repository.go -destination mock/pg_repository.go -package mock
package user

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
)

type PostgresRepository interface {
	Create(user *domain.User) (int, error)
	List() ([]domain.User, error)
	GetByID(userID domain.UserID) (domain.User, error)
	Update(user *domain.User) (domain.User, error)
	Delete(userID domain.UserID) error
}
