package database

import "database/sql"

func (db *appdbimpl) GetMembers(searchname string) (*sql.Rows, error) {

	row, err := db.c.Query("SELECT name FROM users WHERE name = searchname%", searchname)

	return row, err
}
