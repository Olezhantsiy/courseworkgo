package main

import "database/sql"

func getDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bookstorage")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
