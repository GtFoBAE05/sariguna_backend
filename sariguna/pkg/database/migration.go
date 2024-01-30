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

	/*
		crud
		catalog: id, url gambar, catrgori id,deskripsi, status, created_at, updated at
		category: id, nama

		sejarah, visi, misi.
		id tipe teks
		  nama perusaahaan sari guna
		  sejarah dadaw
		  visi semangat
		  visi juang

		https://wa.me/628XXXXXXXXXX?text=Saya%20interested%20in%20your%20car%20for%20sale


	*/

	_, err = db.Exec(query)

	return
}
