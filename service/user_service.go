package service

import (
	"context"
	"net/http"

	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
	"github.com/khilmi-aminudin/dvdrentalv1/models/web"
	"github.com/khilmi-aminudin/dvdrentalv1/repository"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4"
)

type UserService interface {
	Create(ctx context.Context, request web.RequestCreateUser) web.ResponseWeb
	FindAll(ctx context.Context) web.ResponseWeb
	FindByUsername(ctx context.Context, username string) web.ResponseWeb
	Update(ctx context.Context, request web.RequestUpdateUser) web.ResponseWeb
	Delete(ctx context.Context, userid int64) web.ResponseWeb
}

type userService struct {
	Repository repository.UserRepository
	DBConn     *pgx.Conn
	Validator  *validator.Validate
}

func NewuserService(repository repository.UserRepository, dbconn *pgx.Conn, validator *validator.Validate) UserService {
	return &userService{
		Repository: repository,
		Validator:  validator,
		DBConn:     dbconn,
	}
}

func (service *userService) Create(ctx context.Context, request web.RequestCreateUser) web.ResponseWeb {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer tx.Commit(ctx)

	hashedPassword := helper.NewSHA256([]byte(request.Passowrd))

	user := service.Repository.Create(ctx, tx, entity.Users{
		Username: request.Username,
		Passowrd: hashedPassword,
	})

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   user,
	}
}

func (service *userService) FindAll(ctx context.Context) web.ResponseWeb {
	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer tx.Commit(ctx)

	users := service.Repository.FindAll(ctx, tx)

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   users,
	}
}

func (service *userService) FindByUsername(ctx context.Context, username string) web.ResponseWeb {
	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)
	defer tx.Commit(ctx)

	user := service.Repository.FindByUsername(ctx, tx, username)

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   user,
	}
}

func (service *userService) Update(ctx context.Context, request web.RequestUpdateUser) web.ResponseWeb {
	panic("")
}
func (service *userService) Delete(ctx context.Context, userid int64) web.ResponseWeb {
	panic("")
}
