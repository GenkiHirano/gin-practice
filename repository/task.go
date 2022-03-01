package repository

import (
	"github.com/GenkiHirano/gin-practice/model"
)

// Create userの作成
func Create(id int, name, comment string) (*model.User, error) {
	user, err := model.NewUser(id, name, comment)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		ID:      user.ID,
		Name:    user.Name,
		Comment: user.Comment,
	}

	return newUser, nil
}
