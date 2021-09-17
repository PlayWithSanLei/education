package service

import (
	"github.com/impact-eintr/education/internal/dao/db"
	"github.com/impact-eintr/education/internal/inErr"
	"github.com/impact-eintr/education/internal/model"
)

func UpdateUser(reqUser *model.UserReq, id int64) (err error) {
	user := new(model.User)
	user.UserID = id
	user.Username = reqUser.Username
	if reqUser.Role == "user" && reqUser.Unit == "交科" {
		return inErr.ErrInvalidUnit
	}
	user.Unit = reqUser.Unit
	user.Role = reqUser.Role
	user.Password, err = db.UserPassword(id)
	if err != nil {
		return
	}

	err = db.UpdateUsers(user)
	if err != nil {
		return
	}
	return nil

}

func GetUsers() ([]model.UserResp, error) {
	return db.GetUsers()
}

func DeleteUser(id int64) error {
	tmpUser := new(model.User)
	tmpUser.UserID = id

	err := db.DeleteUsersHandler(tmpUser)
	if err != nil {
		return err
	}
	return nil

}
