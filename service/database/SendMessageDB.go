package database

import (
	"log"
	"strconv"
)

func (db *appdbimpl) SendMessage(message Message) (Chat, error) {
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

	_, err = db.c.Exec("INSERT INTO chat_message VALUES (?,?)", message.ChatId, message.Id)
	if err != nil {
		return chat, err
	}

	log.Println("SENDER ID: ", message.Sender.Id)
	_, err = db.c.Exec("INSERT INTO message_user VALUES (?,?)", message.Id, message.Sender.Id)
	if err != nil {
		return chat, err
	}

	chat, err = db.GetConversation(strconv.Itoa(message.ChatId))
	if err != nil {
		return chat, err
	}

	return chat, nil
}
