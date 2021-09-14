package model

type ParamSignUp struct {
	UserName   string `eorm:"username" json:"username" binding:"required"`
	Email      string `eorm:"email" json:"email" binding:"required"`
	Mobile     string `eorm:"mobile" json:"mobile" binding:"required"`
	Password   string `eorm:"password" json:"password" binding:"required"`
	RePassword string `eorm:"repassword" json:"repassword" binding:"required,eqfield=Password"`
	Question   string `eorm:"question" json:"question" binding:"required"`
	Answer     string `eorm:"answer" json:"answer" binding:"required"`
}

type ParamLogin struct {
	Mobile   string `eorm:"mobile" json:"mobile" binding:"required"`
	Password string `eorm:"password" json:"password" binding:"required"`
}

type UserReq struct {
	UserID   string `json:"userid" binding:"required"`
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Unit     string `json:"unit" binding:"required"`
}
