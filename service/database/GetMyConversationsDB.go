package database

func (db *appdbimpl) GetMyConversations(userId string) ([]Chat, error) {

	rowsChat, err := db.c.Query(`SELECT chats.chatId, chats.chatName, chats.picture  
			FROM chats
			JOIN chat_user ON chat_user.chatId = chats.chatId
			WHERE userId = ?`, userId)
	if err != nil {
		return nil, err
	}

	var chats []Chat
	for rowsChat.Next() {
		var chat Chat
		if rowsChat.Scan(&chat.Id, &chat.Name, &chat.Picture) != nil {
			return nil, err
		}

		rowsMembers, err := db.c.Query(
			`SELECT chat_user.userId, users.userName, users.picture
			FROM chats
			JOIN chat_user ON chat_user.chatId = chats.chatId
			JOIN users ON chat_user.userId = users.userId
			WHERE chats.chatId = ?`, chat.Id)
		if err != nil {
			return nil, err
		}
		rowsMessages, err := db.c.Query(
			`SELECT	chat_message.messageId, messages.date, messages.content, messages.comment 
			FROM chats
			JOIN chat_message ON chat_message.chatId = chats.chatId
			JOIN messages ON chat_message.messageId = messages.messageId
			WHERE chats.chatId = ?`, chat.Id)
		if err != nil {
			return nil, err
		}

		for rowsMembers.Next() {
			var member User

			err := rowsMembers.Scan(&member.Id, &member.Name, &member.Picture)
			if err != nil {
				return nil, err
			}
			chat.Members = append(chat.Members, member)
		}

		for rowsMessages.Next() {
			var message Message

			err := rowsMessages.Scan(&message.Id, &message.Date, &message.Content, &message.Comment)
			if err != nil {
				return nil, err
			}

			chat.Messages = append(chat.Messages, message)
		}
		err = rowsMembers.Close()
		if err != nil {
			return chats, err
		}

		err = rowsMessages.Close()
		if err != nil {
			return chats, err
		}
		chats = append(chats, chat)
	}
	err = rowsChat.Close()
	if err != nil {
		return chats, err
	}

	return chats, err
}
