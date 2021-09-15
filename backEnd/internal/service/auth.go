package service

import (
	"fmt"
	"log"

	"github.com/impact-eintr/education/internal/dao/db"
	"github.com/impact-eintr/education/internal/model"
	"github.com/impact-eintr/education/pkg/jwt"
	sf "github.com/impact-eintr/education/pkg/snowflake"
)

func SignUp(p *model.ParamSignUp) error {
	// 检查用户是否已经注册
	if err := db.CheckUserExist(p.Mobile); err != nil {
		return err
	}

	// 生成UID
	userID := sf.GenID()

	// 构造一个User实例
	user := &model.User{
		UserID:   userID,
		Username: p.UserName,
		Password: p.Password,
		Mobile:   p.Mobile,
		Email:    p.Email,
		Question: p.Question,
		Answer:   p.Answer,
		Role:     "user",
		Unit:     "暂无",
	}

	// 存入数据库
	return db.InsertUser(user)

}

// 处理用户登录以及JWT的发放 id name role token
func Login(p *model.ParamLogin) (string, string, string, string, error) {
	// 构造一个User实例
	user := &model.User{
		Mobile:   p.Mobile,
		Password: p.Password,
	}

	// 数据库验证
	if err := db.UserLogin(user); err != nil {
		log.Println(user)
		return "", "", "", "", nil
	}

	// 验证通过后发放token
	aToken, err := jwt.GenToken(user.UserID, user.Username, user.Role)
	return fmt.Sprintf("%d", user.UserID), user.Username, user.Role, aToken, err

}
