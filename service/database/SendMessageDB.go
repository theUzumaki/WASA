package database

import (
	"strconv"
)

func (db *appdbimpl) SendMessage(message Message, userid string, chatid string) (Chat, error) {
	var chat Chat

	var stringId string
	row := db.c.QueryRow("SELECT MAX(messageId) FROM messages")
	err := row.Scan(&stringId)
	if stringId == "" {
		stringId = "0"
	} else if err != nil {
		return chat, err
	}
	message.Id, err = strconv.Atoi(stringId)
	if err != nil {
		return chat, err
	}
	message.Id++

	_, err = db.c.Exec("INSERT INTO messages VALUES (?,?,?,?)", message.Id, message.Date, message.Content, message.Comment)
	if err != nil {
		return chat, err
	}

	_, err = db.c.Exec("INSERT INTO chat_message VALUES (?,?)", chatid, message.Id)
	if err != nil {
		return chat, err
	}

	_, err = db.c.Exec("INSERT INTO message_user VALUES (?,?)", message.Id, userid)
	if err != nil {
		return chat, err
	}

	chat, err = db.GetConversation(chatid)
	if err != nil {
		return chat, err
	}

	return chat, nil
}
