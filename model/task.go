package model

import (
	"errors"
)

type User struct {
	ID      int
	Name    string
	Comment string
}

// NewUser userを初期化
func NewUser(id int, name, Comment string) (*User, error) {
	if id == 0 {
		return nil, errors.New("idを入力してください")
	}

	if name == "" {
		return nil, errors.New("nameを入力してください")
	}

	user := &User{
		ID:      id,
		Name:    name,
		Comment: Comment,
	}

	return user, nil
}
