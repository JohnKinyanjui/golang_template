package db

import (
	"context"
	"fmt"
	query "golang_template/internal/db/generated"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	PgConn *pgxpool.Pool
	once   sync.Once
	Query  *query.Queries
)

func Init() {
	once.Do(func() {
		config, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse database config: %v\n", err)
			os.Exit(1)
		}

		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
			os.Exit(1)
		}

		PgConn = pool
		Query = query.New(pool)

		// Graceful shutdown: Close pool on application exit
		go func() {
			<-context.Background().Done()
			PgConn.Close()
		}()
	})
}
