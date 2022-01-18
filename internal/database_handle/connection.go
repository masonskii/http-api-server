package database_handle

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connection ...
type Connection struct {
	driver        string `toml:"db_driver"`
	database_name string `toml:"db_name`
	username      string `toml:"db_username"`
	password      string `toml:"db_password"`
}

// NewConnection ...
func NewConnection() *Connection {
	return &Connection{
		driver:        "mysql",
		database_name: "FabProjects",
		username:      "root",
		password:      "root",
	}
}

// OpenConnection ...
func (handle *DatabaseHandle) OpenConnection() error {
	conn := NewConnection()
	db, err := sql.Open(conn.driver, "%v:%v@/%v", conn.username, conn.password, conn.database_name)
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()
	if err := db.Ping(); err != nil {
		return err
	}

	handle.db = db
	return nil

}

// CloseConnection ...
func (handle *DatabaseHandle) CloseConnection() error {
	handle.db.Close()
}
