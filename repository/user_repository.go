package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"github.com/khilmi-aminudin/dvdrentalv1/models/entity"
	"github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4"
)

type UserRepository interface {
	Create(ctx context.Context, tx pgx.Tx, user entity.Users) entity.Users
	Update(ctx context.Context, tx pgx.Tx, user entity.Users) entity.Users
	Delete(ctx context.Context, tx pgx.Tx, user entity.Users) error
	FindById(ctx context.Context, tx pgx.Tx, user entity.Users) entity.Users
	FindAll(ctx context.Context, tx pgx.Tx) []entity.Users
	FindByUsername(ctx context.Context, tx pgx.Tx, username string) entity.Users
	NewOTP(ctx context.Context, tx pgx.Tx, username string) string
	ClearOTP(ctx context.Context, tx pgx.Tx, username string, tokens string) bool
}

type userRepository struct{}

func NewuserRepository() UserRepository {
	return &userRepository{}
}

func (repository *userRepository) Create(ctx context.Context, tx pgx.Tx, user entity.Users) entity.Users {
	queryString := "INSERT INTO users(username, password) VALUES ($1,$2);"
	cmdTag, err := tx.Exec(ctx, queryString, user.Username, user.Passowrd)
	helper.LogError(err)
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
	helper.LogError(err)

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
	helper.LogError(err)

	return user
}

func (repository *userRepository) Update(ctx context.Context, tx pgx.Tx, user entity.Users) entity.Users {
	queryString := "UPDATE users SET username = $1, password = $2, last_update $3 WHERE username = $1;"

	cmdTag, err := tx.Exec(ctx, queryString, user.Username, user.Passowrd, user.LastUpdate)
	helper.LogError(err)

	var result entity.Users
	if cmdTag.Update() {
		return result
	}

	result.UserId = user.UserId
	result.Username = user.Username
	result.Passowrd = user.Passowrd
	result.LastUpdate = time.Now()

	return result

}

func (repository *userRepository) Delete(ctx context.Context, tx pgx.Tx, user entity.Users) error {
	queryString := "DELETE FROM actor WHERE user_id = $1;"

	cmdTag, err := tx.Exec(ctx, queryString, user.UserId)

	if !cmdTag.Delete() {
		return err
	}
	return nil
}

func (repository *userRepository) FindById(ctx context.Context, tx pgx.Tx, user entity.Users) entity.Users {
	queryString := "SELECT user_id, username,password,last_upadte FROM users WHERE user_id = $1;"

	row := tx.QueryRow(ctx, queryString, user.UserId)

	var result entity.Users

	err := row.Scan(&result.UserId, &result.Username, &result.Passowrd, &result.LastUpdate)
	// helper.LogError(helper.LoggerInit(),err)
	if err == sql.ErrNoRows {
		return result
	}

	return result
}

func (repository *userRepository) NewOTP(ctx context.Context, tx pgx.Tx, username string) string {
	queryString := "UPDATE users SET tokens = $2 WHERE username = $1 RETURNING tokens;"

	token := helper.EncodeToString(6)

	cmdTag, err := tx.Exec(ctx, queryString, username, token)
	if err != nil {
		logrus.New().Error(err.Error())
		return ""
	}

	if cmdTag.Update() {
		return token
	}
	return ""
}

func (repository *userRepository) ClearOTP(ctx context.Context, tx pgx.Tx, username string, tokens string) bool {
	queryString := "UPDATE users SET tokens = '--' WHERE username = $1 AND tokens = $2 RETURNING tokens;"

	cmdTag, err := tx.Exec(ctx, queryString, username, tokens)

	if err != nil {
		logrus.New().Error(err.Error())
		return false
	}

	if cmdTag.Update() {
		return true
	}

	return false

}
