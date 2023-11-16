package postgres

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type databaseConnector interface {
	Get(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}
