package service

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
	"github.com/khilmi-aminudin/dvdrentalv1/models/web"
	"github.com/khilmi-aminudin/dvdrentalv1/repository"
)

type CategoryService interface {
	Create(ctx context.Context, request web.RequestCreateCategory) web.ResponseWeb
	Update(ctx context.Context, request web.RequestUpdateCategory) web.ResponseWeb
	Delete(ctx context.Context, categoryId int64) web.ResponseWeb
	FindById(ctx context.Context, categoryId int64) web.ResponseWeb
	FindAll(ctx context.Context) web.ResponseWeb
}

type categoryService struct {
	Repository repository.CategoryRepository
	DBConn     *pgx.Conn
	Validator  *validator.Validate
}

func NewCategoryService(repo repository.CategoryRepository, dbconn *pgx.Conn, validator *validator.Validate) CategoryService {
	return &categoryService{
		Repository: repo,
		DBConn:     dbconn,
		Validator:  validator,
	}
}

func (service *categoryService) Create(ctx context.Context, request web.RequestCreateCategory) web.ResponseWeb {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	category := service.Repository.Create(ctx, tx, entity.Category{Name: request.Name})

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   category,
	}

}

func (service *categoryService) Update(ctx context.Context, request web.RequestUpdateCategory) web.ResponseWeb {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	category := service.Repository.Update(ctx, tx, entity.Category{CategoryId: request.CategoryId, Name: request.Name})

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   category,
	}

}

func (service *categoryService) Delete(ctx context.Context, categoryId int64) web.ResponseWeb {
	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	err = service.Repository.Delete(ctx, tx, entity.Category{CategoryId: categoryId})
	if err != nil {
		return web.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		}
	}
	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   err,
	}
}

func (service *categoryService) FindById(ctx context.Context, categoryId int64) web.ResponseWeb {
	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	category := service.Repository.FindById(ctx, tx, entity.Category{CategoryId: categoryId})
	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   category,
	}
}

func (service *categoryService) FindAll(ctx context.Context) web.ResponseWeb {
	tx, err := service.DBConn.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.CommirOrRollback(tx, ctx)

	categories := service.Repository.FindAll(ctx, tx)

	return web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   categories,
	}
}
