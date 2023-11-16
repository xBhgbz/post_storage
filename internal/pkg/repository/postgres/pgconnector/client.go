package pgconnector

import (
	"context"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DatabaseConfig struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Name     string `env:"DB_NAME"`
	SSLMode  string `env:"DB_SSL"`
}

func ConnectToPostgresDB(ctx context.Context) (*PostgresDB, error) {
	var config DatabaseConfig
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.Connect(ctx, generateDSN(&config))
	if err != nil {
		return nil, err
	}
	return &PostgresDB{cluster: pool}, nil
}

func generateDSN(db *DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host,
		db.Port,
		db.User,
		db.Password,
		db.Name,
		db.SSLMode)
}
