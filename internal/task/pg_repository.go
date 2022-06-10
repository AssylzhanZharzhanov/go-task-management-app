package task

import "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"

type PostgresRepository interface {
	Create(task domain.Task) (domain.Task, error)
	List() ([]domain.Task, error)
	GetByID(taskID domain.TaskID) (domain.Task, error)
	Update(task domain.Task) (domain.Task, error)
	Delete(domain.Task, error) error
}
