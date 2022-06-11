package user

import (
	"testing"

	domain "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"

	"github.com/AssylzhanZharzhanov/task-management-app/internal/user/mock"
	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
)

func TestService_Create(t *testing.T) {
	var (
		validUserDTO = &domain.CreateUserDTO{
			FirstName: "Stephen",
			LastName:  "Strange",
			Email:     "doctor.strange@gmail.com",
			Password:  "strange",
		}
		validUser = &domain.User{
			FirstName: "Stephen",
			LastName:  "Strange",
			Email:     "doctor.strange@gmail.com",
			Password:  "strange",
		}
		validUserResult = domain.User{
			ID:        1,
			FirstName: "Stephen",
			LastName:  "Strange",
			Email:     "doctor.strange@gmail.com",
			Password:  "strange",
		}
	)

	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mock.NewMockPostgresRepository(stubCtrl)

	repoStub.EXPECT().
		Create(validUser).
		Return(int(validUserResult.ID), nil).
		AnyTimes()

	// Setup basic service.
	service := NewService(repoStub)

	// Define tests.
	type arguments struct {
		user *domain.CreateUserDTO
	}
	type result struct {
		userID int
	}
	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Success: create valid user",
			arguments: arguments{
				user: validUserDTO,
			},
			expected: result{
				userID: int(validUserResult.ID),
			},
			expectError: false,
		},
		{
			name: "Fail: user is nil",
			arguments: arguments{
				user: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid first_name",
			arguments: arguments{
				user: &domain.CreateUserDTO{
					LastName: "Strange",
					Email:    "doctor.strange@gmail.com",
					Password: "strange",
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid email",
			arguments: arguments{
				user: &domain.CreateUserDTO{
					FirstName: "Stephen",
					LastName:  "Strange",
					Password:  "strange",
				},
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arguments := test.arguments
			expected := test.expected

			createdUserID, err := service.Create(arguments.user)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					userID: createdUserID,
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
		validUserDTO = &domain.UpdateUserDTO{
			ID:        1,
			FirstName: "Stephen",
			LastName:  "Strange",
			Email:     "doctor.strange@gmail.com",
			Password:  "masterofmystic",
		}
		validUser = &domain.User{
			ID:        1,
			FirstName: "Stephen",
			LastName:  "Strange",
			Email:     "doctor.strange@gmail.com",
			Password:  "masterofmystic",
		}
		validUserResult = &domain.User{
			ID:        1,
			FirstName: "Stephen",
			LastName:  "Strange",
			Email:     "doctor.strange@gmail.com",
			Password:  "masterofmystic",
		}
	)

	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mock.NewMockPostgresRepository(stubCtrl)

	repoStub.EXPECT().
		Update(validUser).
		Return(validUserResult, nil).
		AnyTimes()

	// Setup basic service.
	service := NewService(repoStub)

	// Define tests.
	type arguments struct {
		user *domain.UpdateUserDTO
	}
	type result struct {
		userID *domain.User
	}
	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Success: update valid user",
			arguments: arguments{
				user: validUserDTO,
			},
			expected: result{
				userID: validUserResult,
			},
			expectError: false,
		},
		{
			name: "Fail: user is nil",
			arguments: arguments{
				user: nil,
			},
			expectError: true,
		},
		{
			name: "Fail: invalid id",
			arguments: arguments{
				user: &domain.UpdateUserDTO{
					FirstName: "Stephen",
					LastName:  "Strange",
					Email:     "doctor.strange@gmail.com",
					Password:  "masterofmystic",
				},
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arguments := test.arguments
			expected := test.expected

			updatedUser, err := service.Update(arguments.user)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					userID: updatedUser,
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
		validUserID = domain.UserID(1)
		validUser   = &domain.User{
			ID:        1,
			FirstName: "Stephen",
			LastName:  "Strange",
			Email:     "doctor.strange@gmail.com",
			Password:  "masterofmystic",
		}
	)

	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mock.NewMockPostgresRepository(stubCtrl)

	repoStub.EXPECT().
		GetByID(validUserID).
		Return(validUser, nil).
		AnyTimes()

	// Setup basic service.
	service := NewService(repoStub)

	// Define tests.
	type arguments struct {
		userID domain.UserID
	}
	type result struct {
		user *domain.User
	}
	tests := []struct {
		name        string
		arguments   arguments
		expected    result
		expectError bool
	}{
		{
			name: "Success: get user by id",
			arguments: arguments{
				userID: validUserID,
			},
			expected: result{
				user: validUser,
			},
			expectError: false,
		},
		{
			name: "Fail:invalid user id",
			arguments: arguments{
				userID: 0,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arguments := test.arguments
			expected := test.expected

			user, err := service.GetByID(arguments.userID)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					user: user,
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
		userID = domain.UserID(1)
	)

	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mock.NewMockPostgresRepository(stubCtrl)

	repoStub.EXPECT().
		Delete(userID).
		Return(nil).
		AnyTimes()

	// Setup basic service.
	service := NewService(repoStub)

	// Define tests.
	type arguments struct {
		userID domain.UserID
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
			name: "Success: delete user",
			arguments: arguments{
				userID: userID,
			},
			expected: result{
				res: nil,
			},
			expectError: false,
		},
		{
			name: "Fail: delete user",
			arguments: arguments{
				userID: 0,
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arguments := test.arguments
			expected := test.expected

			err := service.Delete(arguments.userID)
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
