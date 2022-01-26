package web

type RequestCreateCategory struct {
	Name string `validate:"required,min=3" json:"name"`
}

type RequestUpdateCategory struct {
	CategoryId int64  `validate:"required" json:"category_id"`
	Name       string `validate:"required,min=3" json:"name"`
}
