package service

import (
	"fmt"

	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
)

type Authentication interface {
	Login(username string, password string) bool
	// ResetPassword()
}

type authentication struct {
	users []entity.Users
}

func NewAuthentication(listusers []entity.Users) Authentication {
	return &authentication{
		users: listusers,
	}
}
func (auth *authentication) Login(username string, password string) bool {

	for _, user := range auth.users {
		fmt.Printf("username : %s\npassword : %s\n", user.Username, user.Passowrd)
		if user.Username == username && user.Passowrd == password {
			return true
		}
	}
	return false
}
