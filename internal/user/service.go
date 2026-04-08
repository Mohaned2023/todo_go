package user

import (
	"context"
	"database/sql"
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

func GetUser(ctx context.Context, conn *sqlx.DB, email string) *User {
	var user User
	err := conn.GetContext(ctx, &user, `
		SELECT
			*
		FROM users
		WHERE email = $1;
	`, email)
	if err != nil {
		if err == sql.ErrNoRows {
			panic(apperr.Exception{
				Type: apperr.UserNotFound,
				More: nil,
			})
		}
		panic(err)
	}
	return &user
}
