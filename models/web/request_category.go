package web

type RequestCreateCategory struct {
	Name string `validate:"required,min=3" json:"name"`
}

type RequestUpdateCategory struct {
	Name string `validate:"required,min=3" json:"name"`
}
