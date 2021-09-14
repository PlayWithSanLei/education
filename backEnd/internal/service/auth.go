package service

import (
	"github.com/impact-eintr/education/internal/dao/db"
	"github.com/impact-eintr/education/internal/model"
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
	}

	// 存入数据库
	return db.InsertUser(user)

}

// 处理用户登录以及JWT的发放
func Login(p *model.ParamLogin) (int64, string, string, string, error) {

}
