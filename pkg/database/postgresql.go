package database

import (
	"context"
	"sync"

	"codeberg.org/kalsin/UtelBot/pkg/configs"
	"github.com/jackc/pgx/v4/pgxpool"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
)

var (
	once sync.Once
	conn *pgxpool.Pool
	err  error
)

func Connect() error {
	once.Do(func() {
		db := configs.GetDBConfig()
		connectionString := db.ConnectionString()
		conn, err = pgxpool.Connect(context.Background(), connectionString)
		if err != nil {
			log.Error.Fatalf("unable to create a database instance: %v\n", db)
		}
	})
	if err != nil {
		log.Error.Printf("unable to connect to database: %v\n", err)
		return err
	}
	return nil
}

func Close() {
	once = sync.Once{}
	conn.Close()
}