package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var DB *sqlx.DB

const DBFile = "scheduler.db"

var schema = `
CREATE TABLE IF NOT EXISTS scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT '',
	title VARCHAR NOT NULL DEFAULT '',
	comment TEXT NOT NULL DEFAULT '',
	repeat VARCHAR NOT NULL DEFAULT ''
);

CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);
`

func init() {
	if err := Init(); err != nil {
		fmt.Printf("db init error: %v\n", err)
	}
}

func Init() error {
	dbFile := DBFile

	if envFile := os.Getenv("TODO_DBFILE"); envFile != "" {
		dbFile = envFile
	}

	var err error
	DB, err = sqlx.Connect("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("open db error: %v", err)
	}

	if _, err := DB.Exec(schema); err != nil {
		return fmt.Errorf("schema init error: %v", err)
	}

	return nil
}
