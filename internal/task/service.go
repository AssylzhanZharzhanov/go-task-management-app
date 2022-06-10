package task

import "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"

type Service interface {
	Create(dto task.CreateTaskDTO) (task.Task, error)
	List() ([]task.Task, error)
	GetByID(taskID task.TaskID) (task.Task, error)
	Update(dto task.UpdateTaskDTO) (task.Task, error)
	Delete(taskID task.TaskID) error
}
