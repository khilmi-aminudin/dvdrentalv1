package web

type LoginCredential struct {
	Username string `validate:"required" json:"username"`
	Passowrd string `validate:"required,min=8" json:"password"`
}
