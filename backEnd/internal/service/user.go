package service

import (
	"github.com/impact-eintr/education/internal/dao/cache"
	"github.com/impact-eintr/education/internal/dao/db"
	"github.com/impact-eintr/education/internal/model"
)

func UpdateUser(reqUser *model.UserReq, id int64) (err error) {
	user := new(model.User)
	err = db.UserData(id, user)
	if err != nil {
		return
	}
	// 替换信息 除了状态和密码
	user.UserID = id
	user.Username = reqUser.Username
	user.Question = reqUser.Question
	user.Answer = reqUser.Answer
	user.Unit = reqUser.Unit

	if user.Role != reqUser.Role {
		// TODO 检测完角色 更新缓存
		cache.CacheUpdate(reqUser.UserID, reqUser.Role)
	}
	user.Role = reqUser.Role

	err = db.UpdateUsers(user)
	if err != nil {
		return
	}
	return nil

}

func GetUsers() ([]model.UserResp, error) {

	// 更新缓存
	users, err := db.GetUsers()

	for i := range users {
		cache.CacheUpdate(users[i].UserID, users[i].Role)
	}
	return users, err
}

func DeleteUser(id int64) error {
	tmpUser := new(model.User)
	tmpUser.UserID = id

	err := db.DeleteUsers(tmpUser)
	if err != nil {
		return err
	}
	return nil

}

func BlockUser(id int64) (err error) {
	user := new(model.User)
	err = db.UserData(id, user)
	if err != nil {
		return err
	}
	if user.Status == 0 {
		user.Status = 1
	} else if user.Status == 1 {
		user.Status = 0
	}

	err = db.UpdateUsers(user)
	if err != nil {
		return
	}
	return nil

}
