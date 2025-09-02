package dao

import "chat-backed/model"

// 模拟数据库
var users = []model.User{
	{ID: 1, Name: "test1", Password: "123"},
	{ID: 2, Name: "test2", Password: "123"},
}

func FindUserByID(id int) *model.User {
	for _, u := range users {
		if u.ID == id {
			return &u
		}
	}
	return nil
}
