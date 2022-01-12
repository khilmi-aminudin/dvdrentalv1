package web

type RequestUpdateActor struct {
	ActorId   int64  `json:"actor_id" validate:"required"`
	FirstName string `json:"first_name" validate:"required,min=3,max=200"`
	LastName  string `json:"last_name" validate:"required,min=3,max=200"`
}
