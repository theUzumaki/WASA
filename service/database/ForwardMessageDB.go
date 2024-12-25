package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) ForwardMessage(userid string, messageid string, chatid string) error {

	row := db.c.QueryRow("SELECT MAX(messageId) FROM messages")
	var newid int
	row.Scan(&newid)

	row = db.c.QueryRow("SELECT * FROM users WHERE userId = ?;", userid)
	err := row.Scan(nil, nil)
	if errors.Is(err, sql.ErrNoRows) {
		return err
	}

	row = db.c.QueryRow("SELECT * FROM messages WHERE messageId = ?", messageid)
	var message Message
	row.Scan(&message.Id, &message.Date, &message.Content, &message.Comment)

	_, err = db.c.Exec("INSERT INTO messages VALUES (?,?,?,?)", newid+1, message.Date, message.Content, message.Comment)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("INSERT INTO chat_message VALUES (?,?)", chatid, newid+1)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("INSERT INTO message_user VALUES (?,?)", newid+1, userid)
	if err != nil {
		return err
	}

	return err
}
