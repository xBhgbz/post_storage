package pgconnector

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDB struct {
	cluster *pgxpool.Pool
}

func (db PostgresDB) Close() {
	db.cluster.Close()
}

func (db PostgresDB) Get(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db.cluster, dst, query, args...)
}

func (db PostgresDB) Select(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db.cluster, dst, query, args...)
}

func (db PostgresDB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.cluster.Exec(ctx, query, args...)
}

func (db PostgresDB) ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.cluster.QueryRow(ctx, query, args...)
}
