package user

import (
	domain "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
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

	userRepository := NewPostgresRepository(sqlDB)

	mockUser := &domain.User{
		ID:        1,
		FirstName: "Stephen",
		LastName:  "Strange",
		Email:     "doctor.strange@gmail.com",
		Password:  "doctor",
		CreatedAt: 1654941048,
	}

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns).AddRow(mockUser.ID)

	createUserQuery := `INSERT INTO "users" ("first_name","last_name","email","password","created_at","id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`

	mock.ExpectBegin()
	mock.ExpectQuery(createUserQuery).WithArgs(
		mockUser.FirstName,
		mockUser.LastName,
		mockUser.Email,
		mockUser.Password,
		mockUser.CreatedAt,
		mockUser.ID,
	).WillReturnRows(rows)
	mock.ExpectCommit()

	createdUserID, err := userRepository.Create(mockUser)
	require.NoError(t, err)
	require.NotNil(t, createdUserID)
	require.Equal(t, mockUser.ID, domain.UserID(createdUserID))
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

	userRepository := NewPostgresRepository(sqlDB)

	userID := domain.UserID(1)
	mockUser := &domain.User{
		ID:        userID,
		FirstName: "Stephen",
		LastName:  "Strange",
		Email:     "doctor.strange@gmail.com",
		Password:  "doctor",
		CreatedAt: 1654941048,
	}

	columns := []string{"id", "first_name", "last_name", "email", "password", "created_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		&mockUser.ID,
		&mockUser.FirstName,
		&mockUser.LastName,
		&mockUser.Email,
		&mockUser.Password,
		&mockUser.CreatedAt,
	)

	findUserByIDQuery := `SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`

	mock.ExpectQuery(findUserByIDQuery).WithArgs(userID).WillReturnRows(rows)
	mock.ExpectCommit()

	foundUser, err := userRepository.GetByID(userID)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.ID, userID)
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

	userRepository := NewPostgresRepository(sqlDB)

	userID := domain.UserID(1)
	mockUser := &domain.User{
		ID:        userID,
		FirstName: "Tony",
		LastName:  "Stark",
		Email:     "tony.stark@gmail.com",
		Password:  "ironman",
		CreatedAt: 1654941048,
	}

	columns := []string{"id", "first_name", "last_name", "email", "password", "created_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		&mockUser.ID,
		&mockUser.FirstName,
		&mockUser.LastName,
		&mockUser.Email,
		&mockUser.Password,
		&mockUser.CreatedAt,
	)

	updateUserQuery := `UPDATE "users" SET "id"=$1,"first_name"=$2,"last_name"=$3,"email"=$4,"password"=$5,"created_at"=$6 WHERE id = $7 RETURNING *`

	mock.ExpectBegin()
	mock.ExpectQuery(updateUserQuery).WithArgs(
		mockUser.ID,
		mockUser.FirstName,
		mockUser.LastName,
		mockUser.Email,
		mockUser.Password,
		mockUser.CreatedAt,
		userID,
	).WillReturnRows(rows)
	mock.ExpectCommit()

	updatedUser, err := userRepository.Update(mockUser)
	require.NoError(t, err)
	require.NotNil(t, updatedUser)
	require.Equal(t, updatedUser.ID, userID)
}
