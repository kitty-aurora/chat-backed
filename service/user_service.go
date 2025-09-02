package service

import (
	"chat-backed/dao"
	"chat-backed/model"
	"errors"
)

func GetUserByID(id int) (*model.User, error) {
	user := dao.FindUserByID(id)
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
