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

func (d *Database) Initialise(dbType, userName, password, dbName string) {
	d.Db, d.Error = sqlx.Connect(dbType, fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", userName, password, dbName, SSLMODE))
	if d.Error != nil {
		log.Fatalln(d.Error)
	}

	defer d.Db.Close()

	d.Tx = d.Db.MustBegin()
}

var DB = &Database{}
