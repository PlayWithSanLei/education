package model

type User struct {
	UserID   int64  `eorm:"userid"`
	Username string `eorm:"username"`
	Password string `eorm:"password"`
	Mobile   string `eorm:"mobile"`
	Email    string `eorm:"email"`
	Question string `eorm:"question"`
	Answer   string `eorm:"answer"`
	Role     string `eorm:"role"`
	Unit     string `eorm:"unit"`
	Status   int8   `eorm:"status"`
}

type UserResp struct {
	UserID   string `json:"userid"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Unit     string `json:"unit"`
	Status   bool   `json:"status"`
}
