package postgresClient

import (
	"context"
	"errors"
	"project/internal/config"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(ctx context.Context, config *config.Postgres) (*pgxpool.Pool, error) {
	if config == nil {
		return nil, errors.New("postgres config must be not empty")
	}

	cfg := &pgxpool.Config{
		ConnConfig: &pgx.ConnConfig{
			Config: pgconn.Config{
				Host:     config.Host,
				Port:     uint16(config.Port),
				Database: config.Database,
			},
		},
	}

	if config.MinConns != nil {
		cfg.MinConns = int32(*config.MinConns)
	}
	if config.MaxConns != nil {
		cfg.MaxConns = int32(*config.MaxConns)
	}
	if config.MaxConnIdleTime != nil {
		maxConnIdleTime := *config.MaxConnIdleTime
		cfg.MaxConnIdleTime = time.Duration(maxConnIdleTime) * time.Second
	}
	if config.MaxConnLifetime != nil {
		maxConnLifetime := *config.MaxConnLifetime
		cfg.MaxConnLifetime = time.Duration(maxConnLifetime) * time.Second
	}

	return pgxpool.NewWithConfig(ctx, cfg)
}
