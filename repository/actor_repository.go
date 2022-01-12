package repository

import (
	"context"
	"fmt"

	"github.com/khilmi-aminudin/dvdrentalv1/helper"

	"github.com/jackc/pgx/v4"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
)

type ActorRepository interface {
	Create(ctx context.Context, tx pgx.Tx, actor entity.Actor) entity.Actor
	Update(ctx context.Context, tx pgx.Tx, actor entity.Actor) entity.Actor
	Delete(ctx context.Context, tx pgx.Tx, actor entity.Actor)
	FindById(ctx context.Context, tx pgx.Tx, actor entity.Actor) entity.Actor
	FindAll(ctx context.Context, tx pgx.Tx) []entity.Actor
	Search(ctx context.Context, tx pgx.Tx, key string) []entity.Actor
}

type actorRepository struct{}

func NewActorRepository() ActorRepository {
	return &actorRepository{}
}

func (repository *actorRepository) Create(ctx context.Context, tx pgx.Tx, actor entity.Actor) entity.Actor {
	queryString := "INSERT INTO actor(first_name,last_name) VALUES($1,$2) RETURNING *;"

	var result entity.Actor
	cmdTag, err := tx.Exec(ctx, queryString, actor.FirstName, actor.LastName)
	helper.PanicIfError(err)

	if cmdTag.RowsAffected() < 1 {
		return result
	}

	result.FirstName = actor.FirstName
	result.LastName = actor.LastName

	return result
}

func (repository *actorRepository) Update(ctx context.Context, tx pgx.Tx, actor entity.Actor) entity.Actor {
	panic("")
}

func (repository *actorRepository) Delete(ctx context.Context, tx pgx.Tx, actor entity.Actor) {
	panic("")
}

func (repository *actorRepository) FindById(ctx context.Context, tx pgx.Tx, actor entity.Actor) entity.Actor {
	panic("")
}

func (repository *actorRepository) FindAll(ctx context.Context, tx pgx.Tx) []entity.Actor {
	queryString := "SELECT actor_id, first_name,last_name,last_update FROM actor;"

	rows, err := tx.Query(ctx, queryString)
	helper.PanicIfError(err)

	var result []entity.Actor
	for rows.Next() {
		var actor entity.Actor
		err = rows.Scan(&actor.ActorId, &actor.FirstName, &actor.LastName, &actor.LastUpdate)
		helper.PanicIfError(err)
		result = append(result, actor)
	}

	return result
}

func (repository *actorRepository) Search(ctx context.Context, tx pgx.Tx, key string) []entity.Actor {
	queryString := fmt.Sprintf("SELECT actor_id, first_name,last_name,last_update FROM actor WHERE LOWER(first_name) LIKE '%%%s' OR LOWER(first_name) LIKE '%s%%' OR LOWER(first_name) LIKE '%%%s%%' OR LOWER(last_name) LIKE '%%%s' OR LOWER(last_name) LIKE '%s%%' OR LOWER(last_name) LIKE '%%%s%%';;", key, key, key, key, key, key)
	rows, err := tx.Query(ctx, queryString)
	helper.PanicIfError(err)

	var result []entity.Actor
	for rows.Next() {
		var actor entity.Actor
		err = rows.Scan(&actor.ActorId, &actor.FirstName, &actor.LastName, &actor.LastUpdate)
		helper.PanicIfError(err)
		result = append(result, actor)
	}

	return result

}
