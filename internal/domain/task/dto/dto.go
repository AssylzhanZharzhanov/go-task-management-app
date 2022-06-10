package dto

import (
	task "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task/entity"
	user "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user/entity"
)

type CreateTaskDTO struct {
	UserID         user.UserID   `json:"user_id"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	StartDate      int64         `json:"start_date"`
	EndDate        int64         `json:"end_date"`
	ReminderPeriod int64         `json:"reminder_period"`
}


type UpdateTaskDTO struct {
	ID             task.TaskID `json:"id"`
	UserID         user.UserID `json:"user_id"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	StartDate      int64         `json:"start_date"`
	EndDate        int64         `json:"end_date"`
	ReminderPeriod int64         `json:"reminder_period"`
}