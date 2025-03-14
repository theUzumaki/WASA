package database

import (
	"strconv"
	"wasatext/service/structs"
)

func (db *appdbimpl) GetMyConversations(userId string) ([]structs.Chat, error) {

	rowsChat, err := db.c.Query(`SELECT chats.chatId, chats.chatName, chats.picture  
			FROM chats
			JOIN chat_user ON chat_user.chatId = chats.chatId
			WHERE userId = ?`, userId)
	if err != nil {
		return nil, err
	}

	var chats []structs.Chat
	for rowsChat.Next() {
		var chat structs.Chat
		if rowsChat.Err() != nil {
			return nil, err
		}

		if rowsChat.Scan(&chat.Id, &chat.Name, &chat.Picture) != nil {
			return nil, err
		}

		chat, err = db.GetConversation(strconv.Itoa(chat.Id))

		chats = append(chats, chat)
	}
	err = rowsChat.Close()
	if err != nil {
		return chats, err
	}

	return chats, err
}
