package model

import (
	"errors"
	"time"
)

type User struct {
	ID        int
	Name      string
	Comment   string
	CreatedAt time.Time
}

// NewUser userを初期化
func NewUser(id int, name, Comment string) (*User, error) {
	if id == 0 {
		return nil, errors.New("idを入力してください")
	}

	if name == "" {
		return nil, errors.New("nameを入力してください")
	}

	now := time.Now()

	user := &User{
		ID:        id,
		Name:      name,
		Comment:   Comment,
		CreatedAt: now,
	}

	return user, nil
}
