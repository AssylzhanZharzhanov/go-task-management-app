package task

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user/dto"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user/entity"
)

type service interface {
	Create(dto *dto.CreateUserDTO) (int, error)
	Get() ([]*entity.User, error)
	GetByID(userID entity.UserID) (*entity.User, error)
	Update(userDTO dto.UpdateUserDTO) (*entity.User, error)
	Delete(userID entity.UserID) error
}
