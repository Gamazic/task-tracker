package domain

import "errors"

var ErrInvalidUsername = errors.New("invalid username")

type Username string

func NewUsername(username string) (Username, error) {
	if username == "" {
		return Username(""), ErrInvalidUsername
	}
	return Username(username), nil
}

type User struct {
	name Username
}

func NewUser(username string) (User, error) {
	usernameDomain, err := NewUsername(username)
	if err != nil {
		return User{}, err
	}
	return User{
		name: usernameDomain,
	}, nil
}

func (u User) GetUsername() Username {
	return u.name
}
