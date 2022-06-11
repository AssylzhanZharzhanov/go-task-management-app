package task

import (
	domain "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestPostgresRepository_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	taskRepository := NewPostgresRepository(sqlDB)

	mockTask := &domain.Task{
		UserID:         1,
		Title:          "Create microservice",
		Description:    "Create microservice using go-kit",
		StartDate:      1654827132,
		EndDate:        1654935132,
		ReminderPeriod: 86400,
		CreatedAt:      1654941048,
	}

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns).AddRow(&mockTask.ID)

	createTaskQuery := `INSERT INTO "tasks" ("user_id","title","description","start_date","end_date","reminder_period","created_at") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`

	mock.ExpectBegin()
	mock.ExpectQuery(createTaskQuery).WithArgs(
		mockTask.UserID,
		mockTask.Title,
		mockTask.Description,
		mockTask.StartDate,
		mockTask.EndDate,
		mockTask.ReminderPeriod,
		mockTask.CreatedAt,
	).WillReturnRows(rows)
	mock.ExpectCommit()

	createdTask, err := taskRepository.Create(mockTask)
	require.NoError(t, err)
	require.NotNil(t, createdTask)
	require.Equal(t, mockTask, createdTask)
	require.Equal(t, mockTask.ID, createdTask.ID)
}

func TestPostgresRepository_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	taskRepository := NewPostgresRepository(sqlDB)

	taskID := domain.TaskID(1)
	mockTask := &domain.Task{
		ID:             1,
		UserID:         1,
		Title:          "Create microservice",
		Description:    "Create microservice using go-kit",
		StartDate:      1654827132,
		EndDate:        1654935132,
		ReminderPeriod: 86400,
		CreatedAt:      1654941048,
	}

	columns := []string{"id", "user_id", "title", "description", "start_date", "end_date", "reminder_period", "created_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		&mockTask.ID,
		&mockTask.UserID,
		&mockTask.Title,
		&mockTask.Description,
		&mockTask.StartDate,
		&mockTask.EndDate,
		&mockTask.ReminderPeriod,
		&mockTask.CreatedAt,
	)

	updateTaskQuery := `UPDATE "tasks" SET "id"=$1,"user_id"=$2,"title"=$3,"description"=$4,"start_date"=$5,"end_date"=$6,"reminder_period"=$7,"created_at"=$8 WHERE id = $9 RETURNING *`

	mock.ExpectBegin()
	mock.ExpectQuery(updateTaskQuery).WithArgs(
		mockTask.ID,
		mockTask.UserID,
		mockTask.Title,
		mockTask.Description,
		mockTask.StartDate,
		mockTask.EndDate,
		mockTask.ReminderPeriod,
		mockTask.CreatedAt,
		taskID,
	).WillReturnRows(rows)
	mock.ExpectCommit()

	createdTask, err := taskRepository.Update(mockTask)
	require.NoError(t, err)
	require.NotNil(t, createdTask)
	require.Equal(t, mockTask, createdTask)
}

func TestPostgresRepository_GetByID(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	taskRepository := NewPostgresRepository(sqlDB)

	taskID := domain.TaskID(1)
	mockTask := &domain.Task{
		UserID:         1,
		Title:          "Create microservice",
		Description:    "Create microservice using go-kit",
		StartDate:      1654827132,
		EndDate:        1654935132,
		ReminderPeriod: 86400,
		CreatedAt:      1654941048,
	}

	columns := []string{"id", "user_id", "title", "description", "start_date", "end_date", "reminder_period", "created_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		&mockTask.ID,
		&mockTask.UserID,
		&mockTask.Title,
		&mockTask.Description,
		&mockTask.StartDate,
		&mockTask.EndDate,
		&mockTask.ReminderPeriod,
		&mockTask.CreatedAt,
	)

	findTaskByIDQuery := `SELECT * FROM "tasks" WHERE id = $1 ORDER BY "tasks"."id" LIMIT 1`

	mock.ExpectQuery(findTaskByIDQuery).WithArgs(taskID).WillReturnRows(rows)
	mock.ExpectCommit()

	foundTask, err := taskRepository.GetByID(taskID)
	require.NoError(t, err)
	require.NotNil(t, foundTask)
	require.Equal(t, mockTask, foundTask)
}
