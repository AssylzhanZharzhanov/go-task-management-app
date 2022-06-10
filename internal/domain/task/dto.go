package domain

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
)

type CreateTaskDTO struct {
	UserID         domain.UserID `json:"user_id"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	StartDate      int64       `json:"start_date"`
	EndDate        int64       `json:"end_date"`
	ReminderPeriod int64       `json:"reminder_period"`
}

type UpdateTaskDTO struct {
	ID             TaskID      `json:"id"`
	UserID         domain.UserID `json:"user_id"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	StartDate      int64       `json:"start_date"`
	EndDate        int64       `json:"end_date"`
	ReminderPeriod int64       `json:"reminder_period"`
}
