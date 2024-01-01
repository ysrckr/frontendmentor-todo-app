package modals

type User struct {
  Id int      `db:"id" json:"id"`
  Ip string    `db:"ip" json:"ip"`
}