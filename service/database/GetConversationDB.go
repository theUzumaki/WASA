package database

func (db *appdbimpl) GetConversation(chatId string) (Chat, error) {
	var chat Chat

	row := db.c.QueryRow(`SELECT *  
			FROM chats
			WHERE chatId = ?`, chatId)

	if err := row.Scan(&chat.Id, &chat.Name); err != nil {
		return chat, err
	}

	rowsMembers, err := db.c.Query(
		`SELECT chat_user.userId, users.userName
		FROM chats
		JOIN chat_user ON chat_user.chatId = chats.chatId
		JOIN users ON chat_user.userId = users.userId
		WHERE chats.chatId = ?`, chat.Id)
	if err != nil {
		return chat, err
	}
	rowsMessages, err := db.c.Query(
		`SELECT	chat_message.messageId, messages.date, messages.content, messages.comment, users.userId, users.userName 
		FROM chats
		JOIN chat_message ON chat_message.chatId = chats.chatId
		JOIN messages ON chat_message.messageId = messages.messageId
		JOIN message_user ON messages.messageId = message_user.messageId
		JOIN users ON message_user.userId = users.userId
		WHERE chats.chatId = ?`, chat.Id)
	if err != nil {
		return chat, err
	}

	defer rowsMembers.Close()
	for rowsMembers.Next() {
		var member User

		err := rowsMembers.Scan(&member.Id, &member.Name)
		if err != nil {
			return chat, err
		}
		chat.Members = append(chat.Members, member)
	}

	defer rowsMessages.Close()
	for rowsMessages.Next() {
		var message Message

		err := rowsMessages.Scan(&message.Id, &message.Date, &message.Content, &message.Comment, &message.Sender.Id, &message.Sender.Name)
		if err != nil {
			return chat, err
		}
		chat.Messages = append(chat.Messages, message)
	}

	return chat, err
}
