package main

import (
	"go-social-backend/internal/env"
	"go-social-backend/internal/store"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         buildDSN(),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}

func buildDSN() string {
	return env.GetString("DB_DSN",
		"postgres://"+env.GetString("DB_USER", "postgres")+
			":"+env.GetString("DB_PASSWORD", "")+
			"@"+env.GetString("DB_HOST", "localhost")+
			":"+env.GetString("DB_PORT", "5432")+
			"/"+env.GetString("DB_NAME", "gosocial")+
			"?sslmode="+env.GetString("DB_SSLMODE", "disable"))
}
