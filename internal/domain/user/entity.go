package domain

import "errors"

const (
	UsersTableName = "users"
)

type UserID int64

type User struct {
	ID        UserID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	FirstName string `json:"first_name" gorm:"not null;column:first_name"`
	LastName  string `json:"last_name" gorm:"not null;column:last_name"`
	Email     string `json:"email" gorm:"not null;column:email"`
	Password  string `json:"password" gorm:"not null;column:password"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at"`
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
