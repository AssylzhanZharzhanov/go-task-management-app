package user

import (
	domain "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/user"
)

type Service struct {
	repository user.PostgresRepository
}

func (s Service) Create(dto domain.CreateUserDTO) (int, error) {
	return 0, nil
}

func (s Service) List() ([]domain.User, error) {
	return nil, nil
}

func (s Service) GetByID(userID domain.UserID) (domain.User, error) {
	return domain.User{}, nil
}

func (s Service) Update(userDTO domain.UpdateUserDTO) (domain.User, error) {
	return domain.User{}, nil
}

func (s Service) Delete(userID domain.UserID) error {
	return nil
}

func NewService(repository user.PostgresRepository) *Service {
	return &Service{repository: repository}
}
