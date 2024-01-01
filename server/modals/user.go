package modals

type User struct {
  Id int      `db:"id"`
  Ip string    `db:"ip"`
}