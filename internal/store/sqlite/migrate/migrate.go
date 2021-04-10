package migrate

import (
	"database/sql"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"

	// Used to initialize the file source driver.
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate connects to the database, checks if it's the correct schema version already,
// and if not, runs the migrations to the target schema version.
func Migrate(db *sql.DB, version uint) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://data/sql/migrations", // Relative to the root of the binary.
		"sqlite3",
		driver,
	)
	if err != nil {
		return err
	}
	if err := m.Migrate(version); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
