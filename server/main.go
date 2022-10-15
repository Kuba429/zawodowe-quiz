package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"zawodowe-quiz/cmd/server"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
	}
	conn, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	server.Init(conn)

}
