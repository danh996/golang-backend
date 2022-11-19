package models

import (
	"database/sql"
)

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		School: School{},
		User:   User{},
	}
}

type Models struct {
	School School
	User   User
}
