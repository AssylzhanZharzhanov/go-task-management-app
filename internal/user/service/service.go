package user

import (
	"errors"
	domain "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/user"
)

type Service struct {
	repository user.PostgresRepository
}

func (s *Service) Create(dto *domain.CreateUserDTO) (int, error) {
	if dto == nil {
		return 0, errors.New("user input is empty")
	}
	if err := dto.Validate(); err != nil {
		return 0, err
	}

	newUser := domain.NewCreatedUser(dto)
	return s.repository.Create(newUser)
}

func (s *Service) List() ([]domain.User, error) {
	return s.repository.List()
}

func (s *Service) GetByID(userID domain.UserID) (domain.User, error) {
	if userID <= 0 {
		return domain.User{}, errors.New("invalid user id")
	}
	return s.repository.GetByID(userID)
}

func (s *Service) Update(userDTO *domain.UpdateUserDTO) (domain.User, error) {
	if userDTO == nil {
		return domain.User{}, errors.New("user is nil")
	}
	if err := userDTO.Validate(); err != nil {
		return domain.User{}, err
	}

	newUser := domain.NewUpdatedUser(userDTO)
	return s.repository.Update(newUser)
}

func (s *Service) Delete(userID domain.UserID) error {
	if userID <= 0 {
		return errors.New("invalid user id")
	}
	return s.repository.Delete(userID)
}

func NewService(repository user.PostgresRepository) *Service {
	return &Service{repository: repository}
}
