package modals

import (
	_ "github.com/lib/pq"
)


type Todo struct {
  Id int            `db:"id" json:"id"`
  Text string       `db:"text" json:"text"`
  Completed bool    `db:"completed" json:"completed"`
  UserId int        `db:"user_id" json:"user_id"`
}