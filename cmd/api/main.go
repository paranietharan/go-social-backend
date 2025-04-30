package main

import (
	"go-social-backend/internal/db"
	"go-social-backend/internal/env"
	"go-social-backend/internal/store"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file ::", err)
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

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	defer db.Close()
	log.Println("Database connection pool established!")

	if err != nil {
		log.Panic(err)
	}
	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}

func buildDSN() string {
	dsn := env.GetString("DB_DSN",
		"postgres://"+env.GetString("DB_USER", "postgres")+
			":"+env.GetString("DB_PASSWORD", "")+
			"@"+env.GetString("DB_HOST", "localhost")+
			":"+env.GetString("DB_PORT", "5432")+
			"/"+env.GetString("DB_NAME", "gosocial")+
			"?sslmode="+env.GetString("DB_SSLMODE", "disable"))
	log.Println("Generated DSN:", dsn)
	return dsn
}
