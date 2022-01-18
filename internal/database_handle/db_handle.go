package database_handle

import "database/sql"

type DatabaseHandle struct {
	db *sql.DB
}
