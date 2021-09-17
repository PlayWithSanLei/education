package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const TokenExpireDuration = time.Hour * 24

var salt = []byte("github.com/impact-eintr")

// MyClaims自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"userid"`
	UserName string `json:"username"`
	UserRole string `json:"userrole"`
	UserUnit string `json:"unit"`
	jwt.StandardClaims
}

func GenToken(userID int64, userName, userRole, userUnit string) (string, error) {
	c := MyClaims{
		UserID:   userID,
		UserName: userName,
		UserRole: userRole,
		UserUnit: userUnit,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "impact-eintr",                             // 签发人
			IssuedAt:  time.Now().Unix(),                          // 签发时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(salt)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString,
		mc,
		func(token *jwt.Token) (interface{}, error) {
			return salt, nil
		})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
