package modals

import (
	_ "github.com/lib/pq"
)


type Todo struct {
  Id int            `db:"id"`
  Text string       `db:"text"`
  Completed bool    `db:"completed"`
  UserId int        `db:"user_id"`
}