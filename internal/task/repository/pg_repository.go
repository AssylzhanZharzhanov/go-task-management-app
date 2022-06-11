package task

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresRepository struct {
	db *gorm.DB
}

func (r *PostgresRepository) Create(task *domain.Task) (*domain.Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return nil, err
	}
	return task, err
}

func (r *PostgresRepository) List() ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *PostgresRepository) GetByID(taskID domain.TaskID) (*domain.Task, error) {
	var task domain.Task
	err := r.db.Where("id = ?", taskID).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *PostgresRepository) Update(task *domain.Task) (*domain.Task, error) {
	var updatedTask domain.Task
	err := r.db.Model(&updatedTask).Clauses(clause.Returning{}).Where("id = ?", task.ID).Updates(task).Error
	if err != nil {
		return nil, err
	}
	return &updatedTask, nil
}

func (r *PostgresRepository) Delete(taskID domain.TaskID) error {
	err := r.db.Delete(&domain.Task{}, taskID).Error
	if err != nil {
		return err
	}
	return nil
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
