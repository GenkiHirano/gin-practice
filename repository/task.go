package repository

import (
	"errors"

	"github.com/GenkiHirano/gin-practice/model"
)

// Create userの作成
func Create(id int, name, comment string) (*model.User, error) {
	if id == 0 {
		return nil, errors.New("idを入力してください")
	}

	if name == "" {
		return nil, errors.New("nameを入力してください")
	}

	newUser := &model.User{
		ID:      id,
		Name:    name,
		Comment: comment,
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
