package database

import (
	"database/sql"
	"log"

	"github.com/kgthegreat/go-crud-starter/config"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	*sql.DB
}

func NewSqliteDB(dbCfg config.SqliteConfig) (*SqliteDB, error) {
	//DSN := fmt.Sprintf("%s:%s@unix(/tmp/sqlite.sock)/%s?parseTime=true", dbCfg.Username, dbCfg.Password, dbCfg.DatabaseName)
	dataSourceName := dbCfg.DatabaseName
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS topics (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, title TEXT NOT NULL, createdAt DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL)")
	if err != nil {
		log.Panic(err)
	}
	statement.Exec()

	return &SqliteDB{db}, nil
}
