package db

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect(){
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	pool , err := pgxpool.New(context.Background(),url)
	if err!=nil {
		panic(fmt.Sprintf("DB unable to connect %v",err))
	}
	DB = pool
}

func Close(){
	DB.Close()
}