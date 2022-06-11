package task

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func (p PostgresRepository) Create(task *domain.Task) (domain.Task, error) {
	return domain.Task{}, nil
}

func (p PostgresRepository) List() ([]domain.Task, error) {
	return nil, nil
}

func (p PostgresRepository) GetByID(taskID domain.TaskID) (domain.Task, error) {
	return domain.Task{}, nil
}

func (p PostgresRepository) Update(task *domain.Task) (domain.Task, error) {
	return domain.Task{}, nil
}

func (p PostgresRepository) Delete(taskID domain.TaskID) error {
	return nil
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
