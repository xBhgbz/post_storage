//go:build integration

package tests

import (
	"context"
	"fmt"
	"log"
	"storage/internal/pkg/environment"
	"storage/internal/pkg/repository/postgres/pgconnector"
	"strings"
	"sync"
	"testing"
)

var db *testPostgresDB

func init() {
	err := environment.LoadEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err = newTestDB()
	if err != nil {
		log.Fatalf("can't create testdb: %s", err.Error())
	}
}

type testPostgresDB struct {
	db *pgconnector.PostgresDB
	sync.Mutex
}

func newTestDB() (*testPostgresDB, error) {
	database, err := pgconnector.ConnectToPostgresDB(context.Background())
	if err != nil {
		return nil, err
	}
	return &testPostgresDB{db: database}, nil
}

func (d *testPostgresDB) Clear(t *testing.T) {
	t.Helper()
	d.Lock()
	defer d.Unlock()

	var tables []string
	err := d.db.Select(context.Background(), &tables, "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name != 'goose_db_version'")
	if err != nil {
		log.Fatal(err.Error())
	}

	query := fmt.Sprintf("Truncate table %s;", strings.Join(tables, ","))
	if _, err := d.db.Exec(context.Background(), query); err != nil {
		log.Fatal(err.Error())
	}
}
