package user

import user "github.com/AssylzhanZharzhanov/task-management-app/internal/user/repository"

type Service struct {
	repository user.PostgresRepository
}

func NewService(repository user.PostgresRepository) *Service {
	return &Service{repository: repository}
}
