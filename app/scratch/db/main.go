package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/sum28it/garage-service/business/data/dbschema"
	"github.com/sum28it/garage-service/business/sys/database"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	dbConfig := database.Config{
		User: "postgres",
		// Connecting to a local instance of postgres
		Password:     "super",
		Host:         "localhost",
		Name:         "postgres",
		MaxIdleConns: 2,
		MaxOpenConns: 0,
		DisableTLS:   true,
	}

	if err := Migrate(dbConfig); err != nil {
		return err
	}

	if err := Seed(dbConfig); err != nil {
		return err
	}

	return nil
}

// ErrHelp provides context that help was given.
var ErrHelp = errors.New("provided help")

// Migrate creates the schema in the database.
func Migrate(cfg database.Config) error {
	db, err := database.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := dbschema.Migrate(ctx, db); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	fmt.Println("migrations complete")
	return nil
}

// Seed loads test data into the database.
func Seed(cfg database.Config) error {
	db, err := database.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := dbschema.Seed(ctx, db); err != nil {
		return fmt.Errorf("seed database: %w", err)
	}

	fmt.Println("seed data complete")
	return nil
}
