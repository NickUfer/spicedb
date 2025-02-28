package migrations

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const errUnableToInstantiate = "unable to instantiate AlembicPostgresDriver: %w"

const postgresMissingTableErrorCode = "42P01"

// AlembicPostgresDriver implements a schema migration facility for use in
// SpiceDB's Postgres datastore.
//
// It is compatible with the popular Python library, Alembic
type AlembicPostgresDriver struct {
	db *sqlx.DB
}

// NewAlembicPostgresDriver creates a new driver with active connections to the database specified.
func NewAlembicPostgresDriver(url string) (*AlembicPostgresDriver, error) {
	connectStr, err := pq.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf(errUnableToInstantiate, err)
	}

	db, err := sqlx.Connect("postgres", connectStr)
	if err != nil {
		return nil, fmt.Errorf(errUnableToInstantiate, err)
	}

	return &AlembicPostgresDriver{db}, nil
}

// Version returns the version of the schema to which the connected database
// has been migrated.
func (apd *AlembicPostgresDriver) Version() (string, error) {
	var loaded string

	if err := apd.db.QueryRowx("SELECT version_num from alembic_version").Scan(&loaded); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == postgresMissingTableErrorCode {
			return "", nil
		}
		return "", fmt.Errorf("unable to load alembic revision: %w", err)
	}

	return loaded, nil
}

// WriteVersion overwrites the value stored to track the version of the
// database schema.
func (apd *AlembicPostgresDriver) WriteVersion(version, replaced string) error {
	result, err := apd.db.Exec("UPDATE alembic_version SET version_num=$1 WHERE version_num=$2", version, replaced)
	if err != nil {
		return fmt.Errorf("unable to update version row: %w", err)
	}

	updatedCount, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("unable to compute number of rows affected: %w", err)
	}

	if updatedCount != 1 {
		return fmt.Errorf("writing version update affected %d rows, should be 1", updatedCount)
	}

	return nil
}
