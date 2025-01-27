package database

import (
	"database/sql"
	"errors"
	"strconv"
)

func (db *appdbimpl) NewChat(userId string, chat Chat) (int, error) {

	row := db.c.QueryRow("SELECT MAX(chatId) FROM chats")
	var stringId string
	var id int
	err := row.Scan(&stringId)
	if stringId == "" {
		id = 0
	} else if err != nil {
		return -1, err
	}
	chat.Id = id + 1

	for i := 0; i < len(chat.Members); i++ {
		for j := 0; j < len(chat.Members); j++ {
			if i != j && chat.Members[i].Id == chat.Members[j].Id {
				return 0, errors.New("duplicate members")
			}
		}
	}

	for i := 0; i < len(chat.Members); i++ {
		row := db.c.QueryRow("SELECT * FROM users WHERE userId = ?;", chat.Members[i].Id)
		err := row.Scan(nil, nil)
		if errors.Is(err, sql.ErrNoRows) {
			return 0, err
		}
	}

	rows, err := db.c.Query("SELECT chatId FROM chats")
	if err != nil {
		return 0, err
	} else {
		defer rows.Close()
		for rows.Next() {
			var id string
			err = rows.Scan(&id)
			if err != nil {
				return 0, err
			}

			i := 0
			for ; i < len(chat.Members); i++ {
				row = db.c.QueryRow("SELECT * FROM chat_user WHERE userId = ? AND chatId = ?;", userId, id)
				err = row.Scan(nil, nil)
				if errors.Is(err, sql.ErrNoRows) {
					break
				}
			}
			if i == len(chat.Members)-1 {
				idInt, err := strconv.Atoi(id)
				if err != nil {
					return 0, err
				}
				return idInt, nil
			}
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
