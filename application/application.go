package application

import (
	"article/pkg/database/migration"
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DbConnectionType uint32
type DbClient struct {
	SqlAdapter *sql.DB
	OrmAdapter *gorm.DB
}

type Application struct {
	Context         context.Context
	DbClients       map[string]*DbClient
	Logger          *logrus.Logger
	MigrationRunner *migration.Runner
	MigrationFlag   string
	IsMigration     bool
	ServiceMode     string
	ServiceName     string
}
