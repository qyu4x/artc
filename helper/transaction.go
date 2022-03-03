package helper

import "database/sql"

func ServiceArticleTransaction(tx *sql.Tx) {
	defer func() {
		err := recover()
		if err != nil {
			errorRolleback := tx.Rollback()
			PanicIfError(errorRolleback)
			panic(err)
		} else {
			err := tx.Commit()
			PanicIfError(err)
		}
	}()
}
