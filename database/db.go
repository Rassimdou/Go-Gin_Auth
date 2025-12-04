package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Rassimdou/Go-Gin_Auth/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB(cfg *config.DBconfig) error {
	//create connecting string
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	//Parser pool config
	poolConf, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err)
	}

	poolConf.MaxConns = 10
	poolConf.MinConns = 2
	poolConf.MaxConnLifetime = time.Hour
	poolConf.MaxConnIdleTime = 30 * time.Minute

	// create a context with timeout for conection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create a connection pool
	pool, err := pgxpool.NewWithConfig(ctx, poolConf)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err)
	}

	//Ping to test connections
	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("unable to ping database: %w", err)
	}

	//assign to global variable
	DB = pool

	log.Println("Successfully connected to database")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
