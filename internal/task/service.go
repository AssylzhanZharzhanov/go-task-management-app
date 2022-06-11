//go:generate mockgen -source service.go -destination mock/service.go -package mock
package task

import "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"

type Service interface {
	Create(dto *domain.CreateTaskDTO) (*domain.Task, error)
	List() ([]domain.Task, error)
	GetByID(taskID domain.TaskID) (*domain.Task, error)
	Update(dto *domain.UpdateTaskDTO) (*domain.Task, error)
	Delete(taskID domain.TaskID) error
}
