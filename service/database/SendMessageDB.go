package database

import (
	"log"
	"strconv"
	"wasatext/service/structs"
)

func (db *appdbimpl) SendMessage(message structs.Message, userId string) (structs.Chat, error) {
	var chat structs.Chat

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
	log.Default().Println("check 1")
	message.Id++
	_, err = db.c.Exec("INSERT INTO messages VALUES (?,?,?,?)", message.Id, message.Date, message.Content, false)
	if err != nil {
		return chat, err
	}
	log.Default().Println("check 2")
	_, err = db.c.Exec("INSERT INTO chat_message VALUES (?,?)", message.ChatId, message.Id)
	if err != nil {
		return chat, err
	}
	log.Default().Println("check 3")
	_, err = db.c.Exec("INSERT INTO message_user VALUES (?,?)", message.Id, message.Sender.Id)
	if err != nil {
		return chat, err
	}
	log.Default().Println("check 4")
	chat, err = db.GetConversation(strconv.Itoa(message.ChatId), userId)
	if err != nil {
		log.Default().Println("Error getting conversation after sending message: ", err)
		return chat, err
	}
	log.Default().Println("check 5")
	return chat, nil
}
