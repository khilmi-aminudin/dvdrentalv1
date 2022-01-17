package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
	"github.com/khilmi-aminudin/dvdrentalv1/models/web"
	"github.com/khilmi-aminudin/dvdrentalv1/repository"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4"
)

type ActorService interface {
	Create(ctx context.Context, requset web.RequestCreateActor) web.ResponseWeb
	FindAll(ctx context.Context) web.ResponseWeb
	Search(ctx context.Context, key string) web.ResponseWeb
	Update(ctx context.Context, request web.RequestUpdateActor) web.ResponseWeb
	Delete(ctx context.Context, actorId int64) web.ResponseWeb
	FindById(ctx context.Context, actorId int64) web.ResponseWeb
}

type actorService struct {
	Repository repository.ActorRepository
	DBConn     *pgx.Conn
	Validator  *validator.Validate
}

func NewActorService(repository repository.ActorRepository, db *pgx.Conn, validator *validator.Validate) ActorService {
	return &actorService{
		Repository: repository,
		DBConn:     db,
		Validator:  validator,
	}
}

func (service *actorService) Create(ctx context.Context, requset web.RequestCreateActor) web.ResponseWeb {
	err := service.Validator.Struct(requset)
	helper.PanicIfError(err)

	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	actor := service.Repository.Create(ctx, tx, entity.Actor{FirstName: requset.FirstName, LastName: requset.LastName})

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actor,
	}
}

func (service *actorService) FindAll(ctx context.Context) web.ResponseWeb {
	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	actors := service.Repository.FindAll(ctx, tx)

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actors,
	}
}

func (service *actorService) Search(ctx context.Context, key string) web.ResponseWeb {
	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	actors := service.Repository.Search(ctx, tx, key)

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actors,
	}
}

func (service *actorService) Update(ctx context.Context, request web.RequestUpdateActor) web.ResponseWeb {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	actor := service.Repository.Update(ctx, tx, entity.Actor{ActorId: request.ActorId, FirstName: request.FirstName, LastName: request.LastName})

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actor,
	}
}

func (service *actorService) Delete(ctx context.Context, actorId int64) web.ResponseWeb {
	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	actor := service.Repository.FindById(ctx, tx, entity.Actor{ActorId: actorId})

	var emptyActor entity.Actor

	if actor == emptyActor {
		return web.ResponseWeb{
			Code:   http.StatusNotFound,
			Status: "Not Found",
		}
	}
	err = service.Repository.Delete(ctx, tx, entity.Actor{ActorId: actorId})

	if err != nil {
		return web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		}
	}

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   fmt.Sprintf("Actor with id %d was deleted", actorId),
	}
}

func (service *actorService) FindById(ctx context.Context, actorId int64) web.ResponseWeb {
	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	actor := service.Repository.FindById(ctx, tx, entity.Actor{ActorId: actorId})

	var emptyActor entity.Actor
	if actor == emptyActor {
		return web.ResponseWeb{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "Data not found",
		}
	}

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actor,
	}
}
