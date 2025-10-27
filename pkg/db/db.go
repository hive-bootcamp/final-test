package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var (
	db     *sql.DB
	schema = `
		CREATE TABLE IF NOT EXISTS scheduler (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date CHAR(8) NOT NULL DEFAULT '',
			title VARCHAR(256) NOT NULL DEFAULT '',
			comment TEXT DEFAULT '',
			repeat VARCHAR(128) DEFAULT ''
		);

		CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);
		`
)

func Init(dbFile string) error {
    if envPath := os.Getenv("TODO_DBFILE"); envPath != "" {
        dbFile = envPath
    }

    install := false
    if _, err := os.Stat(dbFile); err != nil {
        if os.IsNotExist(err) {
            install = true
        } else {
            return fmt.Errorf("ошибка проверки файла: %w", err)
        }
    }

    var err error
    db, err = sql.Open("sqlite", dbFile)
    if err != nil {
        return fmt.Errorf("ошибка открытия: %w", err)
    }

    if install {
        if _, err := db.Exec(schema); err != nil {
            return fmt.Errorf("ошибка создания таблицы: %w", err)
        }
    }

    return nil
}


func GetDB() *sql.DB {
	return db
}
