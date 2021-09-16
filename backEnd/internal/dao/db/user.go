package db

import (
	"context"
	"database/sql"

	"github.com/impact-eintr/education/global"
	"github.com/impact-eintr/education/internal/inErr"
	"github.com/impact-eintr/education/internal/model"
	"github.com/impact-eintr/eorm"
)

func NoUser() error {
	users := []model.User{}
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").
		Select("*")

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(context.Background(), statement, &users)
	if err == nil && len(users) == 0 {
		return inErr.ErrNoUser
	}

	return err

}

func GetUsers() (users []model.UserResp, err error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").
		Select("userid, username, role, unit")

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err = c.FindAll(context.Background(), statement, &users)
	if err != nil {
		return nil, err
	}
	return

}

var userMap = map[string]string{
	"用户名":  "username",
	"用户id": "userid",
	"单位":   "unit",
	"角色":   "role",
}

func QueryUsers(column string, v interface{}) (users []model.UserResp, err error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user")
	switch column {
	case "userid":
		statement = statement.AndEqual(column, v.(int64))
	default:
		statement = statement.AndLike(column, v.(string))
	}
	statement = statement.Select("userid, username, role, unit")

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err = c.FindAll(context.Background(), statement, &users)
	if err != nil {
		return
	}
	return

}

func UserPassword(id int64) (string, error) {
	tmp := new(model.User)
	statement1 := eorm.NewStatement()
	statement1 = statement1.SetTableName("user").
		AndEqual("userid", id).
		Select("userid, username, password, role, unit")

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindOne(context.Background(), statement1, tmp)
	// 查询没有结果
	if err == sql.ErrNoRows {
		return "", inErr.ErrUserNotExist
	}

	// 查询失败
	if err != nil {
		return "", err
	}
	return tmp.Password, nil

}

func UpdateUsers(users *model.User) (err error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").
		AndEqual("userid", users.UserID).UpdateStruct(users)

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	_, err = c.Update(context.Background(), statement)
	if err != nil {
		return
	}
	return

}

func DeleteUsersHandler(users *model.User) (err error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").
		AndEqual("userid", users.UserID)

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	_, err = c.Delete(context.Background(), statement)
	if err != nil {
		return
	}
	return

}
