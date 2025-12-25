package main

import (
	"log"
	"strings"

	"github.com/sayanmondal31/gosocial/internal/db"
	"github.com/sayanmondal31/gosocial/internal/env"
	"github.com/sayanmondal31/gosocial/internal/store"
)

func main() {
	addr := env.GetString(
		"DB_ADDR",
		"postgres://admin:adminpassword@localhost/social?sslmode=disable",
	)

	// ensure sslmode is set (disable by default)
	if !strings.Contains(addr, "sslmode=") {
		if strings.Contains(addr, "?") {
			addr += "&sslmode=disable"
		} else {
			addr += "?sslmode=disable"
		}
	}

	log.Printf("DB_ADDR=%s\n", addr)

	conn, err := db.New(addr, 30, 30, "15m")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(*store)
}
