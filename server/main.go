package main

import (
	"context"
	"fmt"
	"os"
	"zawodowe-quiz/cmd/server"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	DB_URL := "postgres://localhost:5432/quiz_zawodowe"
	conn, err := pgxpool.New(context.Background(), DB_URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	server.Init(conn)

}
