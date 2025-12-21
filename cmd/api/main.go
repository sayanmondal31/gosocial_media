package main

import (
	"log"

	"github.com/sayanmondal31/gosocial/internal/db"
	"github.com/sayanmondal31/gosocial/internal/env"
	"github.com/sayanmondal31/gosocial/internal/store"
)

const apiversion = "0.0.1"

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: &dbConfig{
			addr: env.GetString(
				"DB_ADDR",
				"postgres://admin:adminpassword@localhost/social?sslmode=disable",
			),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),      //higher the conns , higher the concurrency made
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),      // need to search on theory
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"), // need to search on theory
		},
		env: env.GetString("ENV", "development"),
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
		store:  *store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
