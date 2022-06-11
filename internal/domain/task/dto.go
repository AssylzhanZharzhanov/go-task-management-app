package domain

import (
	"errors"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
)

type CreateTaskDTO struct {
	UserID         domain.UserID `json:"user_id"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	StartDate      int64         `json:"start_date"`
	EndDate        int64         `json:"end_date"`
	ReminderPeriod int64         `json:"reminder_period"`
}

func (t CreateTaskDTO) Validate() error {
	if len(t.Title) == 0 {
		return errors.New("title is invalid")
	}
	return nil
}

type UpdateTaskDTO struct {
	ID             TaskID        `json:"id"`
	UserID         domain.UserID `json:"user_id"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	StartDate      int64         `json:"start_date"`
	EndDate        int64         `json:"end_date"`
	ReminderPeriod int64         `json:"reminder_period"`
}

func (t UpdateTaskDTO) Validate() error {
	if len(t.Title) == 0 {
		return errors.New("title is invalid")
	}
	return nil
}
