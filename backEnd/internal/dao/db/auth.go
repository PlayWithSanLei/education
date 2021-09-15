package db

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/impact-eintr/education/global"
	"github.com/impact-eintr/education/internal/inErr"
	"github.com/impact-eintr/education/internal/model"
	"github.com/impact-eintr/eorm"
)

// 加密salt
const salt string = `impact-eintr`

// 检查注册时用户是否已经存在
func CheckUserExist(mobile string) error {
	sqlStr := `select count(userid) from user where username = ?`
	var count int

	err := global.DB.QueryRow(sqlStr, mobile).Scan(&count)
	if err != nil {
		log.Println(err)
		return err
	}

	if count > 0 {
		return inErr.ErrUserExist
	}
	return nil

}

// 注册一个新的用户
func InsertUser(user *model.User) (err error) {
	// 密码加密
	user.Password = encryptPassword(user.Password)

	// 执行SQL语句入库
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").InsertStruct(user)

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	_, err = c.Insert(context.Background(), statement)
	if err != nil {
		return err
	}

	// 成功将用户注册后 返回用户的UID
	return nil
}

// 加密函数 (md5)
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// 用户登录
func UserLogin(user *model.User) (err error) {
	oPassword := user.Password

	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").
		AndEqual("mobile", user.Mobile).
		Select("*")

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err = c.FindOne(context.Background(), statement, user)
	if err == sql.ErrNoRows {
		return inErr.ErrUserNotExist
	}

	if err != nil {
		// 查询失败
		return err
	}
	// 判断密码是否匹配
	password := encryptPassword(oPassword)
	if password != user.Password {
		log.Println(password, user.Password)
		return inErr.ErrInvalidPassword
	}

	return
}
