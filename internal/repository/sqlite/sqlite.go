package sqlite

import "database/sql"

type Sqlite struct {
	db *sql.DB
}

func newSqlite()