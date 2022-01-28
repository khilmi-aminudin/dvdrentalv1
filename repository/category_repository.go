package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx pgx.Tx, category entity.Category) entity.Category
	Update(ctx context.Context, tx pgx.Tx, category entity.Category) entity.Category
	Delete(ctx context.Context, tx pgx.Tx, category entity.Category) error
	FindById(ctx context.Context, tx pgx.Tx, category entity.Category) entity.Category
	FindAll(ctx context.Context, tx pgx.Tx) []entity.Category
}

type categoryRepository struct{}

func NewcategoryRespository() CategoryRepository {
	return &categoryRepository{}
}
func (repository *categoryRepository) Create(ctx context.Context, tx pgx.Tx, category entity.Category) entity.Category {
	queryString := "INSERT INTO category(name) VALUES($1);"

	cmdTag, err := tx.Exec(ctx, queryString, category.Name)
	helper.LogError(err)
	if cmdTag.Insert() {
		return category
	}
	return entity.Category{}
}

func (repository *categoryRepository) Update(ctx context.Context, tx pgx.Tx, category entity.Category) entity.Category {
	queryString := "UPDATE category SET name = $2 WHERE id = $1 RETURNING category_id, name, last_update;"

	cmdTag, err := tx.Exec(ctx, queryString, category.CategoryId, category.Name)
	helper.LogError(err)

	if cmdTag.Update() {
		category.LastUpdate = time.Now()
		return category
	}
	return entity.Category{}
}

func (repository *categoryRepository) Delete(ctx context.Context, tx pgx.Tx, category entity.Category) error {
	queryString := "DELETE FROM category WHERE category_id=$1;"

	cmdTag, err := tx.Exec(ctx, queryString, category.CategoryId)

	if cmdTag.Delete() {
		return nil
	}
	return err
}

func (repository *categoryRepository) FindById(ctx context.Context, tx pgx.Tx, category entity.Category) entity.Category {
	queryString := "SELECT category_id, name, last_update FROM category WHERE category_id = $1;"
	row := tx.QueryRow(ctx, queryString, category.CategoryId)

	var result entity.Category
	err := row.Scan(&result.CategoryId, &result.Name, &result.LastUpdate)

	if err != sql.ErrNoRows {
		return result
	}

	return result
}

func (repository *categoryRepository) FindAll(ctx context.Context, tx pgx.Tx) []entity.Category {
	queryString := "SELECT category_id, name, last_update FROM category;"

	rows, err := tx.Query(ctx, queryString)
	helper.LogError(err)

	var categories []entity.Category
	for rows.Next() {
		var c entity.Category
		err = rows.Scan(&c.CategoryId, &c.Name, &c.LastUpdate)
		helper.LogError(err)

		categories = append(categories, c)
	}
	return categories
}
