package task

import (
	"errors"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
)

type TaskID int64

type Task struct {
	ID             TaskID      `json:"id"`
	UserID         user.UserID `json:"user_id"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	StartDate      int64       `json:"start_date"`
	EndDate        int64       `json:"end_date"`
	ReminderPeriod int64       `json:"reminder_period"`
	CreatedAt      int64       `json:"created_at"`
}

func NewNote(dto CreateTaskDTO) Task {
	return Task{
		UserID:         dto.UserID,
		Title:          dto.Title,
		Description:    dto.Description,
		StartDate:      dto.StartDate,
		EndDate:        dto.EndDate,
		ReminderPeriod: dto.ReminderPeriod,
	}
}

func NewUpdatedNote(dto UpdateTaskDTO) Task {
	return Task{
		ID:             dto.ID,
		UserID:         dto.UserID,
		Title:          dto.Title,
		Description:    dto.Description,
		StartDate:      dto.StartDate,
		EndDate:        dto.EndDate,
		ReminderPeriod: dto.ReminderPeriod,
	}
}

func (t Task) Validate() error {
	if len(t.Title) == 0 {
		return errors.New("title is invalid")
	}
	return nil
}
