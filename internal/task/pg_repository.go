//go:generate mockgen -source pg_repository.go -destination mock/pg_repository.go -package mock
package task

import "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"

type PostgresRepository interface {
	Create(task *domain.Task) (*domain.Task, error)
	List() ([]domain.Task, error)
	GetByID(taskID domain.TaskID) (*domain.Task, error)
	Update(task *domain.Task) (*domain.Task, error)
	Delete(taskID domain.TaskID) error
	IsTaskExist(userID int64, startDate int64) (bool, error)
}
