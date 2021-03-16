// Package main provides ...
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gtldhawalgandhi/go-training/4.Advanced/api"
	"github.com/gtldhawalgandhi/go-training/4.Advanced/db"
	"github.com/gtldhawalgandhi/go-training/4.Advanced/util"
	"github.com/jackc/pgx/v4/pgxpool"
)

func startServer(config util.Config, store db.Store) {
	var err error
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("Failed to create server", err)
	}

	err = server.Start(":5555")
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	util.Dump(2, config)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgxpool.Connect(ctx, config.DBSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	store := db.NewPGStore(conn)
	startServer(config, store)

}
