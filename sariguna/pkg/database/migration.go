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

		CREATE TABLE IF NOT EXISTS product_category(
			id serial PRIMARY KEY,
			category_name varchar
		);

		CREATE TABLE IF NOT EXISTS product(
			id serial PRIMARY KEY,
			category_id integer,
			name varchar(150),
			description varchar(150),
			status bool,
			image_url varchar(150),
			created_at timestamp DEFAULT NOW(),
			updated_at timestamp DEFAULT NOW()
		);

		ALTER TABLE product DROP CONSTRAINT IF EXISTS fk_product_category;

		ALTER TABLE product
		ADD CONSTRAINT fk_product_category
		FOREIGN KEY (category_id)
		REFERENCES product_category(id);

	`

	/*
		crud
		catalog: id, url gambar, catrgori id,name, deskripsi, status, created_at, updated at
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
