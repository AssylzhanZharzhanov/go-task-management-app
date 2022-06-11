package task

import (
	"errors"
	domain "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/task"
)

type Service struct {
	repository task.PostgresRepository
}

func (s *Service) Create(dto *domain.CreateTaskDTO) (*domain.Task, error) {
	if dto == nil {
		return &domain.Task{}, errors.New("input is invalid")
	}
	if err := dto.Validate(); err != nil {
		return &domain.Task{}, err
	}

	newTask := domain.NewCreatedTask(dto)

	isExist, err := s.repository.IsTaskExist(int64(newTask.UserID), newTask.StartDate)
	if err != nil {
		return &domain.Task{}, err
	}
	if isExist {
		return &domain.Task{}, errors.New("user has task during this time")
	}

	return s.repository.Create(newTask)
}

func (s *Service) List() ([]domain.Task, error) {
	return s.repository.List()
}

func (s *Service) GetByID(taskID domain.TaskID) (*domain.Task, error) {
	if taskID <= 0 {
		return &domain.Task{}, errors.New("invalid task id")
	}
	return s.repository.GetByID(taskID)
}

func (s *Service) Update(dto *domain.UpdateTaskDTO) (*domain.Task, error) {
	if dto == nil {
		return &domain.Task{}, errors.New("input is invalid")
	}
	if err := dto.Validate(); err != nil {
		return &domain.Task{}, err
	}
	newTask := domain.NewUpdatedTask(dto)

	return s.repository.Update(newTask)
}

func (s *Service) Delete(taskID domain.TaskID) error {
	if taskID <= 0 {
		return errors.New("invalid task id")
	}
	return s.repository.Delete(taskID)
}

func NewService(repository task.PostgresRepository) *Service {
	return &Service{repository: repository}
}
