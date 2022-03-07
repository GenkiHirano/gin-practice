package repository

import (
	"errors"

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

// Get userの取得
func Get(id int) (*model.User, error) {
	if id == 1 {
		getUser := &model.User{
			ID:      1,
			Name:    "太郎",
			Comment: "よろしくお願いします",
		}

		return getUser, nil
	}

	return nil, errors.New("ユーザーが見つかりませんでした")
}
