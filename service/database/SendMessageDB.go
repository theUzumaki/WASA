package database

import (
	"strconv"
)

func (db *appdbimpl) SendMessage(message Message, userid string, chatid string) error {

	row := db.c.QueryRow("SELECT MAX(messageId) FROM messages")
	var stringId string
	err := row.Scan(&stringId)
	if stringId == "" {
		stringId = "0"
	} else if err != nil {
		return err
	}
	message.Id, err = strconv.Atoi(stringId)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("INSERT INTO messages VALUES (?,?,?,?)", message.Id+1, message.Date, message.Content, message.Comment)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("INSERT INTO chat_message VALUES (?,?)", chatid, message.Id)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("INSERT INTO message_user VALUES (?,?)", message.Id, userid)
	if err != nil {
		return err
	}

	return err
}
