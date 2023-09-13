package migration

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/sirupsen/logrus"
)

const (
	fileScheme  = "file"
	mysqlScheme = "mysql"
)

type DriverFunc func() (string, database.Driver, error)
type SourceFunc func() (string, source.Driver, error)

type Runner struct {
	logger           *logrus.Logger
	sourceName       string
	databaseName     string
	sourceInstance   source.Driver
	databaseInstance database.Driver
}

func NewRunner(src SourceFunc, db DriverFunc, logger *logrus.Logger) (*Runner, error) {
	sourceName, srcDriver, err := src()
	if err != nil {
		return nil, err
	}
	databaseName, dbDriver, err := db()
	if err != nil {
		return nil, err
	}
	return &Runner{
		sourceName:       sourceName,
		databaseName:     databaseName,
		sourceInstance:   srcDriver,
		databaseInstance: dbDriver,
		logger:           logger,
	}, nil
}

func (r *Runner) Up() error {
	fields := logrus.Fields{
		"source_type":   r.sourceName,
		"database_type": r.databaseName,
	}
	migration, err := migrate.NewWithInstance(
		r.sourceName,
		r.sourceInstance,
		r.databaseName,
		r.databaseInstance,
	)
	if err != nil {
		r.logger.WithFields(fields).Error(err)
		return err
	}
	defer func() {
		r.logger.WithFields(fields).Info("Finish Migration")
		migration.Close()
	}()

	r.logger.WithFields(fields).Info("Starting Migration...")
	err = migration.Up()
	if err == migrate.ErrNoChange {
		r.logger.WithFields(fields).Info(err)
		return nil
	}
	if err != nil {
		r.logger.WithFields(fields).Error(err)
	}
	return err
}

func (r *Runner) Rollback() error {
	fields := logrus.Fields{
		"source_type":   r.sourceName,
		"database_type": r.databaseName,
	}
	migration, err := migrate.NewWithInstance(
		r.sourceName,
		r.sourceInstance,
		r.databaseName,
		r.databaseInstance,
	)
	if err != nil {
		r.logger.WithFields(fields).Error(err)
		return err
	}
	defer func() {
		r.logger.WithFields(fields).Info("Finish Rollback")
		migration.Close()
	}()
	r.logger.WithFields(fields).Info("Starting Rollback...")
	err = migration.Down()
	if err == migrate.ErrNoChange {
		r.logger.WithFields(fields).Info(err)
		return nil
	}
	if err != nil {
		r.logger.WithFields(fields).Error(err)
	}
	return err
}
