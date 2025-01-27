package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) LeaveGroup(chatId string, userId string) error {

	_, err := db.c.Exec("DELETE FROM chat_user WHERE chatId = ? AND userId = ?", chatId, userId)
	if err != nil {
		return err
	}

	row := db.c.QueryRow("SELECT chatId FROM chat_user WHERE chatId = ?", chatId)

	if errors.Is(row.Scan(nil), sql.ErrNoRows) {
		_, err := db.c.Exec("DELETE FROM chats WHERE chatId = ?", chatId)
		if err != nil {
			return err
		}
	}

	return err
}
