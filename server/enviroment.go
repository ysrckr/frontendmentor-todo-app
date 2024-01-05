package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Enviroment struct {
	Mode       string
	PORT       string
	DB         string
	DBUSER     string
	DBPASSWORD string
	DBNAME     string
  ALLOWEDORIGINS []string
}

func (e *Enviroment) SetEnviromentVars() {
	if e.Mode == "development" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Println(err)
		}
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Println(err)
		}
	}

	e.PORT = os.Getenv("API_PORT")
	e.DB = os.Getenv("DB_DRIVER")
	e.DBUSER = os.Getenv("DB_USER")
	e.DBPASSWORD = os.Getenv("DB_PASSWORD")
	e.DBNAME = os.Getenv("DB_NAME")
  e.ALLOWEDORIGINS = strings.Split(os.Getenv("ALLOWED_ORIGINS"), ";")
}

func (e *Enviroment) SetEnviromentMode(mode string) {
	if mode != "" {
		e.Mode = mode
	} else {
		e.Mode = "production"
	}
}

var Env = &Enviroment{}
