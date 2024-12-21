package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) NewChat(userId string, chat Chat) error {

	row := db.c.QueryRow("SELECT MAX(chatId) FROM chats")
	var id int
	row.Scan(&id)
	chat.Id = id + 1

	for i := 0; i < len(chat.Members); i++ {
		row := db.c.QueryRow("SELECT * FROM users WHERE userId = ?;", chat.Members[i].Id)
		err := row.Scan(nil, nil)
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	_, err := db.c.Exec("INSERT INTO chats VALUES (?, ?);", chat.Id, chat.Name)
	if err != nil {
		return err
	}

	for i := 0; i < len(chat.Members); i++ {
		_, err = db.c.Exec("INSERT INTO chat_user VALUES (?, ?);", chat.Id, chat.Members[i].Id)
		if err != nil {
			return err
		}
	}
	return nil
}
