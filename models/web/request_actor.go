package web

type RequestCreateActor struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=200"`
	LastName  string `json:"last_name" validate:"required,min=3,max=200"`
}
