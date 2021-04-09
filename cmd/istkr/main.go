package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	issuectrl "github.com/amckinney/issue-tracker/internal/controller/issue"
	"github.com/amckinney/issue-tracker/internal/handler/issue"
	"github.com/amckinney/issue-tracker/internal/persist"
	"github.com/amckinney/issue-tracker/internal/persist/sqlite"
	"github.com/amckinney/issue-tracker/internal/persist/sqlite/migrate"
	"github.com/amckinney/issue-tracker/internal/router"
	"go.uber.org/zap"
)

const (
	_localhost              = "127.0.0.1"
	_defaultPort            = 3000
	_inMemorySource         = ":memory:"
	_migrationSchemaVersion = 1
)

var _defaultAddress = fmt.Sprintf("%s:%d", _localhost, _defaultPort)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		exit(os.Stderr, err)
	}
	db, err := sqlite.New(
		&sqlite.Config{
			Source: "data/sql/db.sql",
		},
	)
	if err != nil {
		exit(os.Stderr, err)
	}
	if err := migrate.Migrate(db, _migrationSchemaVersion); err != nil {
		exit(os.Stderr, err)
	}
	store, err := persist.NewSQL(db)
	if err != nil {
		exit(os.Stderr, err)
	}
	defer store.Close()
	router := router.New(
		issue.New(
			logger,
			store,
			issuectrl.New(),
		),
	)
	logger.Info("Server successfully started")
	exit(os.Stderr, http.ListenAndServe(_defaultAddress, router))
}

// exit writes out the error, and exits with a
// non-zero code if the error is non-nil.
func exit(writer io.Writer, err error) {
	if err == nil {
		os.Exit(0)
	}
	io.WriteString(writer, err.Error())
	os.Exit(1)
}
