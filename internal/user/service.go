package user

import (
	"errors"
	"todo/internal/apperr"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// - Throws UserFound
// - Throws any, Database error.
func UserCreate(conn *sqlx.DB, user *User) *User {
	var insertedUser User;
	d, err := conn.NamedQuery(`
		INSERT INTO users (name, age, email, password)
		VALUES (:name, :age, :email, :password)
		RETURNING *;
	`, user);
	if err == nil && d.Rows.Next() {
		d.StructScan(&insertedUser);
		d.Close();
		return &insertedUser;
	}
	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			panic(apperr.Exception{
				Type: apperr.UserFound,
				More: nil,
			})
		}
	}
	panic(err)
}
