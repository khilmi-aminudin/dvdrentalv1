package helper

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func CommirOrRollback(tx pgx.Tx, ctx context.Context) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback(ctx)
		if errorRollback != nil {
			LoggerInit().Warn(errorRollback.Error())
		}

	} else {
		errorCommit := tx.Commit(ctx)
		if errorCommit != nil {
			LoggerInit().Warn(errorCommit.Error())
		}
	}
}
