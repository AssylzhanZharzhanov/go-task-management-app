package user

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func (p PostgresRepository) Create(user domain.User) (int, error) {
	return 0, nil
}

func (p PostgresRepository) List() ([]domain.User, error) {
	return nil, nil
}

func (p PostgresRepository) GetByID(userID domain.UserID) (domain.User, error) {
	return domain.User{}, nil
}

func (p PostgresRepository) Update(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (p PostgresRepository) Delete(userID domain.UserID) error {
	return nil
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
