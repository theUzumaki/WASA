package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) NewChat(userId string, chat Chat) (int, error) {

	row := db.c.QueryRow("SELECT MAX(chatId) FROM chats")
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	chat.Id = id + 1

	for i := 0; i < len(chat.Members); i++ {
		row := db.c.QueryRow("SELECT * FROM users WHERE userId = ?;", chat.Members[i].Id)
		err := row.Scan(nil, nil)
		if errors.Is(err, sql.ErrNoRows) {
			return 0, err
		}
	}

	_, err = db.c.Exec("INSERT INTO chats VALUES (?, ?);", chat.Id, chat.Name)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(chat.Members); i++ {
		_, err = db.c.Exec("INSERT INTO chat_user VALUES (?, ?);", chat.Id, chat.Members[i].Id)
		if err != nil {
			return 0, err
		}
	}
	return chat.Id, nil
}
