package database

import "github.com/jmoiron/sqlx"

func Migration(db *sqlx.DB) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id uuid PRIMARY KEY,
			fullname varchar(150),
			email varchar(150),
			password varchar(150),
			role varchar(150),
			created_at timestamp DEFAULT NOW(),
			updated_at timestamp DEFAULT NOW()
		);
	`

	_, err = db.Exec(query)

	return
}
