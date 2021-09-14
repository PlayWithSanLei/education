package model

type User struct {
	UserID   int64  `eorm:"userid"`
	Username string `eorm:"username"`
	Password string `eorm:"password"`
	Mobile   string `eorm:"mobile"`
	Email    string `eorm:"email"`
	Role     string `eorm:"role"`
	Question string `eorm:"question"`
	Answer   string `eorm:"answer"`
}

type UserResp struct {
	UserID   string `json:"userid"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Unit     string `json:"unit"`
}
