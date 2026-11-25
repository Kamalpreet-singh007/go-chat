package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/Kamalpreet-singh007/go-chat/internals/server"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
	db, err := sql.Open("pgx", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Error opening database", err.Error())
	}
	db.Ping()
	log.Println("Database connected")
	server.StartServer()
}
