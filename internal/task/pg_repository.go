package task

import "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"

type PostgresRepository interface {
	Create(task task.Task) (task.Task, error)
	List() ([]task.Task, error)
	GetByID(taskID task.TaskID) (task.Task, error)
	Update(task task.Task) (task.Task, error)
	Delete(task.Task, error) error
}
