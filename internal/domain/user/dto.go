package domain

import "errors"

type CreateUserDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u CreateUserDTO) Validate() error {
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

type UpdateUserDTO struct {
	ID        UserID `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u UpdateUserDTO) Validate() error {
	if u.ID <= 0 {
		return errors.New("user id is invalid")
	}
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

func NewCreatedUser(dto *CreateUserDTO) *User {
	return &User{
		FirstName: dto.FirstName,
		LastName: dto.LastName,
		Email: dto.Email,
		Password: dto.Password,
	}
}

func NewUpdatedUser(dto *UpdateUserDTO) *User {
	return &User{
		ID: dto.ID,
		FirstName: dto.FirstName,
		LastName: dto.LastName,
		Email: dto.Email,
		Password: dto.Password,
	}
}
