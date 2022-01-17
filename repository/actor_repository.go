package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/khilmi-aminudin/dvdrentalv1/helper"

	"github.com/jackc/pgx/v4"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
)

type ActorRepository interface {
	Create(ctx context.Context, tx pgx.Tx, actor entity.Actor) entity.Actor
	Update(ctx context.Context, tx pgx.Tx, actor entity.Actor) entity.Actor
	Delete(ctx context.Context, tx pgx.Tx, actor entity.Actor) error
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
	queryString := "UPDATE actor SET first_name = $2, last_name = $3 WHERE actor_id = $1;"

	cmdTag, err := tx.Exec(ctx, queryString, actor.ActorId, actor.FirstName, actor.LastName)
	helper.PanicIfError(err)

	var result entity.Actor
	if cmdTag.RowsAffected() < 1 {
		return result
	}

	result.ActorId = actor.ActorId
	result.FirstName = actor.FirstName
	result.LastName = actor.LastName
	result.LastUpdate = time.Now()

	return result
}

func (repository *actorRepository) Delete(ctx context.Context, tx pgx.Tx, actor entity.Actor) error {
	queryString := "DELETE FROM actor WHERE actor_id = $1"

	cmdTag, err := tx.Exec(ctx, queryString, actor.ActorId)

	if !cmdTag.Delete() {
		return err
	}
	return nil
}

func (repository *actorRepository) FindById(ctx context.Context, tx pgx.Tx, actor entity.Actor) entity.Actor {
	queryString := "SELECT actor_id, first_name, last_name, last_update FROM actor WHERE actor_id = $1"

	row := tx.QueryRow(ctx, queryString, actor.ActorId)

	var result entity.Actor

	err := row.Scan(&result.ActorId, &result.FirstName, &result.LastName, &result.LastUpdate)
	// helper.PanicIfError(err)
	if err == *&sql.ErrNoRows {
		return result
	}

	return result

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
