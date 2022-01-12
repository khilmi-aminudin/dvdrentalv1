package helper

import "database/sql"

func CommirOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		if errorRollback != nil {
			panic(errorRollback.Error())
		}

	} else {
		errorCommit := tx.Commit()
		if errorCommit != nil {
			panic(errorCommit.Error())
		}
	}
}
