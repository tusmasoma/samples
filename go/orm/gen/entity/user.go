package entity

import "errors"

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func NewUser(id, name, email, password string) (*User, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
