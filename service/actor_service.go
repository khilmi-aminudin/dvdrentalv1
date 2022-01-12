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

type ActorService interface {
	Create(ctx context.Context, requset web.RequestCreateActor) web.ResponseWeb
	FindAll(ctx context.Context) web.ResponseWeb
	Search(ctx context.Context, key string) web.ResponseWeb
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

	defer tx.Commit(ctx)

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

	defer tx.Commit(ctx)

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

	defer tx.Commit(ctx)

	actors := service.Repository.Search(ctx, tx, key)

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actors,
	}
}
