package main

import (
	"github.com/CP-Payne/exercise/internal/env"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":3000"),
		env:  env.GetString("ENV", "development"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5432/exercisedb?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	app := NewApp(&cfg)
	app.run()

}
