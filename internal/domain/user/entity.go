package user

import "errors"

const (
	UserTableName = "users"
)

type UserID int64

type User struct {
	ID        UserID `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

func (u User) Validate() error {
	if len(u.FirstName) == 0 {
		return errors.New("first name is invalid")
	}
	if len(u.LastName) == 0 {
		return errors.New("last name is invalid")
	}
	if len(u.Email) == 0 {
		return errors.New("email is invalid")
	}
	if len(u.Password) == 0 {
		return errors.New("password is invalid")
	}

	return nil
}
