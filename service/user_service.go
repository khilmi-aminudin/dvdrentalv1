package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
	"github.com/khilmi-aminudin/dvdrentalv1/models/web"
	"github.com/khilmi-aminudin/dvdrentalv1/repository"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4"
)

type UserService interface {
	Create(ctx context.Context, request web.RequestCreateUser) web.ResponseWeb
	Update(ctx context.Context, request web.RequestUpdateUser) web.ResponseWeb
	Delete(ctx context.Context, userid int64) web.ResponseWeb
	FindById(ctx context.Context, userid int64) web.ResponseWeb
	FindAll(ctx context.Context) web.ResponseWeb
	FindByUsername(ctx context.Context, username string) web.ResponseWeb
	NewOTP(ctx context.Context, username string) web.ResponseWeb
	ClearOTP(ctx context.Context, username string, tokens string) web.ResponseWeb
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

	conn := service.DBConn
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	hashedPassword := helper.NewSHA256([]byte(request.Passowrd))

	user := service.Repository.Create(ctx, tx, entity.Users{
		Email:    request.Email,
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
	conn := service.DBConn
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	users := service.Repository.FindAll(ctx, tx)

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   users,
	}
}

func (service *userService) FindByUsername(ctx context.Context, username string) web.ResponseWeb {
	conn := service.DBConn
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx, ctx)

	user := service.Repository.FindByUsername(ctx, tx, username)

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   user,
	}
}

func (service *userService) Update(ctx context.Context, request web.RequestUpdateUser) web.ResponseWeb {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	conn := service.DBConn
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	user := service.Repository.Update(ctx, tx, entity.Users{
		Username:   request.Username,
		Passowrd:   request.Passowrd,
		LastUpdate: time.Now(),
	})

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   user,
	}

}

func (service *userService) Delete(ctx context.Context, userid int64) web.ResponseWeb {
	conn := service.DBConn
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	err = service.Repository.Delete(ctx, tx, entity.Users{UserId: userid})

	if err != nil {
		return web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data: gin.H{
				"message": err.Error(),
			},
		}
	}
	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
	}
}

func (service *userService) FindById(ctx context.Context, userid int64) web.ResponseWeb {
	conn := service.DBConn
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	user := service.Repository.FindById(ctx, tx, entity.Users{UserId: userid})

	var emptyUser entity.Users
	if user == emptyUser {
		return web.ResponseWeb{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data: gin.H{
				"message": fmt.Sprintf("user with id %d not found", userid),
			},
		}
	}
	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   user,
	}
}

func (service *userService) NewOTP(ctx context.Context, username string) web.ResponseWeb {
	conn := service.DBConn
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	otp := service.Repository.NewOTP(ctx, tx, username)
	if otp == "" {
		return web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
		}
	}

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   otp,
	}
}

func (service *userService) ClearOTP(ctx context.Context, username string, tokens string) web.ResponseWeb {
	conn := service.DBConn
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	cleared := service.Repository.ClearOTP(ctx, tx, username, tokens)

	if cleared {
		return web.ResponseWeb{
			Code:   http.StatusOK,
			Status: "Success",
		}
	}

	return web.ResponseWeb{
		Code:   http.StatusBadRequest,
		Status: "Bad Request",
	}
}
