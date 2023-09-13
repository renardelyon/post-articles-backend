package gorm

import (
	"article/config"
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlORM(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s",
		cfg.ArticleDatabase.User,
		cfg.ArticleDatabase.Password,
		cfg.ArticleDatabase.Protocol,
		cfg.ArticleDatabase.Host,
		cfg.ArticleDatabase.Port,
		cfg.ArticleDatabase.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func PingDB(db *sql.DB) error {
	_, err := db.Exec("SELECT 1")
	return err
}
