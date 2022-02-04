package web

type RequestCreateUser struct {
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required,min=3" json:"username"`
	Passowrd string `validate:"required,min=8" json:"password"`
}

type RequestUpdateUser struct {
	UserId   int64  `json:"user_id"`
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required,min=3" json:"username"`
	Passowrd string `validate:"required,min=8" json:"password"`
}

type RequestUpdateToken struct {
	Tokens string `validate:"required" json:"tokens"`
}
