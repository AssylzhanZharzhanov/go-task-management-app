package user

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresRepository struct {
	db *gorm.DB
}

func (r *PostgresRepository) Create(user *domain.User) (int, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return int(user.ID), nil
}

func (r *PostgresRepository) List() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *PostgresRepository) GetByID(userID domain.UserID) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return &domain.User{}, err
	}
	return &user, nil
}

func (r *PostgresRepository) Update(user *domain.User) (*domain.User, error) {
	var updatedUser domain.User
	err := r.db.Model(&updatedUser).Clauses(clause.Returning{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return &domain.User{}, err
	}
	return &updatedUser, nil
}

func (r *PostgresRepository) Delete(userID domain.UserID) error {
	err := r.db.Delete(&domain.User{}, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
