package repository

import (
	"context"

	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"

	"github.com/jackc/pgx/v4"
)

type UserRepository interface {
	Create(ctx context.Context, tx pgx.Tx, user entity.Users) entity.Users
	FindAll(ctx context.Context, tx pgx.Tx) []entity.Users
	FindByUsername(ctx context.Context, tx pgx.Tx, username string) entity.Users
}

type userRepository struct{}

func NewuserRepository() UserRepository {
	return &userRepository{}
}

func (repository *userRepository) Create(ctx context.Context, tx pgx.Tx, user entity.Users) entity.Users {
	queryString := "INSERT INTO users(username, password) VALUES ($1,$2);"
	cmdTag, err := tx.Exec(ctx, queryString, user.Username, user.Passowrd)
	helper.PanicIfError(err)
	if cmdTag.Insert() {
		return entity.Users{
			Username: user.Username,
			Passowrd: user.Passowrd,
		}
	}
	return entity.Users{}
}

func (repository *userRepository) FindAll(ctx context.Context, tx pgx.Tx) []entity.Users {
	queryString := "SELECT user_id, username ,password, last_update FROM users;"

	rows, err := tx.Query(ctx, queryString)
	helper.PanicIfError(err)

	var users []entity.Users

	for rows.Next() {
		var i entity.Users
		rows.Scan(
			&i.UserId,
			&i.Username,
			&i.Passowrd,
			&i.LastUpdate,
		)

		users = append(users, i)
	}

	return users
}

func (repository *userRepository) FindByUsername(ctx context.Context, tx pgx.Tx, username string) entity.Users {
	queryString := "SELECT user_id, username ,password, last_update FROM users WHERE username = $1;"

	row := tx.QueryRow(ctx, queryString, username)

	var user entity.Users
	err := row.Scan(&user.UserId, &user.Username, &user.Passowrd, &user.LastUpdate)
	helper.PanicIfError(err)

	return user
}
