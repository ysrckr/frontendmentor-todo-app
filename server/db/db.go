package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const SSLMODE = "disabled"

type Database struct {
	Db    *sqlx.DB
	Error error
}

func (d *Database) InitilizeDatabase(dbType, userName, dbName string) *sqlx.Tx {
	d.Db, d.Error = sqlx.Connect(dbType, fmt.Sprintf("user=%s dbname=%s sslmode=%s", userName, dbName, SSLMODE))
	if d.Error != nil {
		log.Fatalln(d.Error)
	}

	tx := d.Db.MustBegin()

	return tx
}
