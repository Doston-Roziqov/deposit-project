package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Doston-Roziqov/deposit-project/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	cfg := config.NewConfig(".")
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)

	dbPool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()
	fmt.Println("Connect Successfully")
}
