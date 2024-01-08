package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const SSLMODE = "disable"

type Database struct {
	Db    *sqlx.DB
	Error error
	Tx    *sqlx.Tx
}

func (d *Database) Initialise(dbType, userName, password, dbName, hostname string) {
	d.Db, d.Error = sqlx.Connect(dbType, fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s", userName, password, dbName, SSLMODE, hostname))
	if d.Error != nil {
		log.Fatalln(d.Error)
	}

	d.Tx = d.Db.MustBegin()
}

var DB = &Database{}
