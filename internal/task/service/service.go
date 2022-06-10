package task

import (
	domain "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/task"
)

type Service struct {
	repository task.PostgresRepository
}

func (s Service) Create(dto domain.CreateTaskDTO) (domain.Task, error) {
	return domain.Task{}, nil
}

func (s Service) List() ([]domain.Task, error) {
	return nil, nil
}

func (s Service) GetByID(taskID domain.TaskID) (domain.Task, error) {
	return domain.Task{}, nil
}

func (s Service) Update(dto domain.UpdateTaskDTO) (domain.Task, error) {
	return domain.Task{}, nil
}

func (s Service) Delete(taskID domain.TaskID) error {
	return nil
}

func NewService(repository task.PostgresRepository) *Service {
	return &Service{repository: repository}
}