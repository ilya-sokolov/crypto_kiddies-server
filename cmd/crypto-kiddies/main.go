package main

import (
	"github.com/ilya-sokolov/crypto_kiddies-server/database"
	"github.com/ilya-sokolov/crypto_kiddies-server/model"
	"github.com/ilya-sokolov/crypto_kiddies-server/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	userName := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	_, err = database.Connect(dbHost, dbName, userName, password)
	model.InitMigration()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = database.Disconnect(database.DB)
		if err != nil {
			panic(err)
		}
	}()
	app := routes.Router()
	err = app.Run(":4000")
	if err != nil {
		panic(err)
	}
}
