package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	PgConn *pgxpool.Pool
	once   sync.Once
)

func init() {
	once.Do(func() {
		config, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse database config: %v\n", err)
			os.Exit(1)
		}

		// Optionally, you can configure additional settings on the pool here:
		// Example: config.MaxConns = 10

		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
			os.Exit(1)
		}

		PgConn = pool

		// Graceful shutdown: Close pool on application exit
		go func() {
			<-context.Background().Done()
			PgConn.Close()
		}()
	})
}
