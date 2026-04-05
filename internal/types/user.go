package types

import "time"

type User struct {
	Id        int64     `db:"id"         json:"id"`
	Name      string    `db:"name"       json:"name"`
	Age       int       `db:"age"        json:"age"`
	Email     string    `db:"email"      json:"email"`
	Password  string    `db:"password"   json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
