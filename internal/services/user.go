package services

import (
	"github.com/jmoiron/sqlx"

	"todo/internal/types"
)

func UserCreate(conn *sqlx.DB, user *types.User) (*types.User, error) {
	var insertedUser types.User;
	d, err := conn.NamedQuery(`
		INSERT INTO users (name, age, email, password)
		VALUES (:name, :age, :email, :password)
		RETURNING *;
	`, user);
	if err == nil && d.Rows.Next() {
		d.StructScan(&insertedUser);
		d.Close();
		return &insertedUser, nil;
	}
	return nil, err;
}
