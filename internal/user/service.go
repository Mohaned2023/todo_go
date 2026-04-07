package user

import "github.com/jmoiron/sqlx"

func UserCreate(conn *sqlx.DB, user *User) (*User, error) {
	var insertedUser User;
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
