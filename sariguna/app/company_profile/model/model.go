package model

type CompanyProfile struct {
	Id      int    `db:"id"`
	Sejarah string `db:"sejarah"`
	Visi    string `db:"visi"`
	Misi    string `db:"misi"`
}
