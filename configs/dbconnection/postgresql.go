package dbconnection

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Path string = "./configs/env/db.env"

func DbConnection() *sql.DB {
	EnvLoader(Path)
	psqlInfo := os.Getenv("POSTGRESQL")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("postgresql doesn't work: %s", err)
	}
	fmt.Println("connected to db successfully!")
	return db
}
func EnvLoader(pathToEnv string) {
	err := godotenv.Load(pathToEnv)
	if err != nil {
		log.Fatalf("Some error occured.Err: %s", err)
	}
}
