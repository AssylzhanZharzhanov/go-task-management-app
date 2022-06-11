package task

import (
	domain "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/task/mock"
	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestService_Create(t *testing.T) {
	var (
		validTaskDTO = &domain.CreateTaskDTO{
			UserID:      1,
			Title:       "Create microservice",
			Description: "Create microservice using go-kit",
			StartDate:   1654827132,
			EndDate:     1654935132,
		}
		validTask = &domain.Task{
			UserID:      1,
			Title:       "Create microservice",
			Description: "Create microservice using go-kit",
			StartDate:   1654827132,
			EndDate:     1654935132,
		}
		validTaskResult = domain.Task{
			ID:          1,
			UserID:      1,
			Title:       "Create microservice",
			Description: "Create microservice using go-kit",
			StartDate:   1654827132,
			EndDate:     1654935132,
		}
	)

	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mock.NewMockPostgresRepository(stubCtrl)

	repoStub.EXPECT().
		Create(validTask).
		Return(validTaskResult, nil).
		AnyTimes()

	// Setup basic service.
	service := NewService(repoStub)

	// Define tests.
	type arguments struct {
		task *domain.CreateTaskDTO
	}
	type result struct {
		task domain.Task
	}
	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Success: create valid task",
			arguments: arguments{
				task: validTaskDTO,
			},
			expected: result{
				task: validTaskResult,
			},
			expectError: false,
		},
		{
			name: "Fail: task is nil",
			arguments: arguments{
				task: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: task title nil",
			arguments: arguments{
				task: nil,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arguments := test.arguments
			expected := test.expected

			createdTask, err := service.Create(arguments.task)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					task: createdTask,
				}
				if diff := deep.Equal(expected, actual); diff != nil {
					t.Error(diff)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestService_Update(t *testing.T) {
	var (
		validTaskDTO = &domain.UpdateTaskDTO{
			ID:          1,
			UserID:      1,
			Title:       "Update microservice",
			Description: "Update microservice using go-micro",
			StartDate:   1654827132,
			EndDate:     1654935132,
		}
		validUpdatedTask = &domain.Task{
			ID:          1,
			UserID:      1,
			Title:       "Update microservice",
			Description: "Update microservice using go-micro",
			StartDate:   1654827132,
			EndDate:     1654935132,
		}
		validTaskResult = domain.Task{
			ID:          1,
			UserID:      1,
			Title:       "Update microservice",
			Description: "Update microservice using go-micro",
			StartDate:   1654827132,
			EndDate:     1654935132,
		}
	)

	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mock.NewMockPostgresRepository(stubCtrl)

	repoStub.EXPECT().
		Update(validUpdatedTask).
		Return(validTaskResult, nil).
		AnyTimes()

	// Setup basic service.
	service := NewService(repoStub)

	// Define tests.
	type arguments struct {
		task *domain.UpdateTaskDTO
	}
	type result struct {
		task domain.Task
	}
	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Success: create valid task",
			arguments: arguments{
				task: validTaskDTO,
			},
			expected: result{
				task: validTaskResult,
			},
			expectError: false,
		},
		{
			name: "Fail: task is nil",
			arguments: arguments{
				task: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: task title nil",
			arguments: arguments{
				task: nil,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arguments := test.arguments
			expected := test.expected

			createdTask, err := service.Update(arguments.task)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					task: createdTask,
				}
				if diff := deep.Equal(expected, actual); diff != nil {
					t.Error(diff)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestService_GetByID(t *testing.T) {
	var (
		validTaskID = domain.TaskID(1)
		validTask   = domain.Task{
			ID:          1,
			UserID:      1,
			Title:       "Create microservice",
			Description: "Create microservice using go-kit",
			StartDate:   1654827132,
			EndDate:     1654935132,
		}
	)

	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mock.NewMockPostgresRepository(stubCtrl)

	repoStub.EXPECT().
		GetByID(validTaskID).
		Return(validTask, nil).
		AnyTimes()

	// Setup basic service.
	service := NewService(repoStub)

	// Define tests.
	type arguments struct {
		taskID domain.TaskID
	}
	type result struct {
		task domain.Task
	}
	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Success: get task by id",
			arguments: arguments{
				taskID: validTaskID,
			},
			expected: result{
				task: validTask,
			},
			expectError: false,
		},
		{
			name: "Fail: task id is invalid",
			arguments: arguments{
				taskID: 0,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arguments := test.arguments
			expected := test.expected

			user, err := service.GetByID(arguments.taskID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					task: user,
				}
				if diff := deep.Equal(expected, actual); diff != nil {
					t.Error(diff)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}

}

func TestService_Delete(t *testing.T) {
	var (
		taskID = domain.TaskID(1)
	)

	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mock.NewMockPostgresRepository(stubCtrl)

	repoStub.EXPECT().
		Delete(taskID).
		Return(nil).
		AnyTimes()

	// Setup basic service.
	service := NewService(repoStub)

	// Define tests.
	type arguments struct {
		taskID domain.TaskID
	}
	type result struct {
		res error
	}
	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Success: delete task",
			arguments: arguments{
				taskID: taskID,
			},
			expected: result{
				res: nil,
			},
			expectError: false,
		},
		{
			name: "Fail: task id is invalid",
			arguments: arguments{
				taskID: 0,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arguments := test.arguments
			expected := test.expected

			err := service.Delete(arguments.taskID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					res: err,
				}
				if diff := deep.Equal(expected, actual); diff != nil {
					t.Error(diff)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}

}
