package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) CommentMessage(messageid string, comment string) error {

	row := db.c.QueryRow("SELECT * FROM messages WHERE messageId = ?", messageid)
	if err := row.Scan(nil, nil, nil, nil); errors.Is(err, sql.ErrNoRows) {
		return err
	}

	_, err := db.c.Exec("UPDATE messages SET comment = ? WHERE messageId = ?", comment, messageid)
	if err != nil {
		return err
	}

	return err
}
