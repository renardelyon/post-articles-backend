package application

import (
	"article/pkg/database/migration"
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type DbConnectionType uint32
type DbClient struct {
	Type       DbConnectionType
	SqlAdapter *sql.DB
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
