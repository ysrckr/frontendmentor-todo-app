package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Enviroment struct {
	Mode       string
	PORT       string
	DB         string
	DBUSER     string
	DBPASSWORD string
	DBNAME     string
}

func (e *Enviroment) SetEnviromentVars() {
	if e.Mode == "development" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Println(err)
		}
	}

	e.PORT = os.Getenv("API_PORT")
	e.DB = os.Getenv("DB")
	e.DBUSER = os.Getenv("DB_USER")
	e.DBPASSWORD = os.Getenv("DB_PASSWORD")
	e.DBNAME = os.Getenv("DB_NAME")
}

var Env = &Enviroment{
	Mode: "development",
}
