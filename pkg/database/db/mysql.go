package db

import (
	"article/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s",
		cfg.ArticleDatabase.User,
		cfg.ArticleDatabase.Password,
		cfg.ArticleDatabase.Protocol,
		cfg.ArticleDatabase.Host,
		cfg.ArticleDatabase.Port,
		cfg.ArticleDatabase.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = PingDB(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func PingDB(db *sql.DB) error {
	_, err := db.Exec("SELECT 1")
	return err
}
