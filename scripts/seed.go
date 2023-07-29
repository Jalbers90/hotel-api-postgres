package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Jalbers90/hotel-api-postgres/types"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

func main() {
	config := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "admin",
		Password: "admin",
		Database: "hotel",
		SSLMode:  "disable",
	}
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.SSLMode,
	)
	// Open a PostgreSQL database.
	// pgconn := pgdriver.NewConnector(
	// 	pgdriver.WithNetwork("tcp"),
	// 	pgdriver.WithAddr("localhost:5432"),
	// 	// pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
	// 	pgdriver.WithUser("admin"),
	// 	pgdriver.WithPassword("admin"),
	// 	pgdriver.WithDatabase("hotel"),
	// 	pgdriver.WithApplicationName("DB:hotel"),
	// 	pgdriver.WithTimeout(5*time.Second),
	// 	pgdriver.WithDialTimeout(5*time.Second),
	// 	pgdriver.WithReadTimeout(5*time.Second),
	// 	pgdriver.WithWriteTimeout(5*time.Second),
	// 	pgdriver.WithConnParams(map[string]interface{}{
	// 		"search_path": "my_search_path",
	// 	}),
	// )
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connStr)))
	db := bun.NewDB(pgdb, pgdialect.New())
	// defer db.Close()

	// Create users table.
	// _, err := db.NewCreateTable().Model((*types.User)(nil)).Exec(context.Background())
	// if err != nil {
	// 	fmt.Println("hi there")
	// 	log.Fatal(err)
	// }

	user := &types.User{
		FirstName: "Amy",
		LastName:  "Albers",
	}
	_, err := db.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
