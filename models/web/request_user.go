package web

type RequestCreateUser struct {
	Username string `validate:"required,min=3" json:"username"`
	Passowrd string `validate:"required,min=8" json:"password"`
}

type RequestUpdateUser struct {
	UserId   int64  `json:"user_id"`
	Username string `validate:"required,min=3" json:"username"`
	Passowrd string `validate:"required,min=8" json:"password"`
}
