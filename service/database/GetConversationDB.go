package database

func (db *appdbimpl) GetConversation(chatId string) (Chat, error) {
	var chat Chat

	row := db.c.QueryRow(`SELECT *  
			FROM chats
			WHERE chatId = ?`, chatId)

	if err := row.Scan(&chat.Id, &chat.Name, &chat.Picture); err != nil {
		return chat, err
	}

	rowsMembers, err := db.c.Query(
		`SELECT chat_user.userId, users.userName, users.picture
		FROM chats
		JOIN chat_user ON chat_user.chatId = chats.chatId
		JOIN users ON chat_user.userId = users.userId
		WHERE chats.chatId = ?`, chat.Id)
	if err != nil {
		return chat, err
	}
	rowsMessages, err := db.c.Query(
		`SELECT	chat_message.messageId, messages.date, messages.content, users.userId, users.userName, users.picture, chats.chatId 
		FROM chats
		JOIN chat_message ON chat_message.chatId = chats.chatId
		JOIN messages ON chat_message.messageId = messages.messageId
		JOIN message_user ON messages.messageId = message_user.messageId
		JOIN users ON message_user.userId = users.userId
		WHERE chats.chatId = ?`, chat.Id)
	if err != nil {
		return chat, err
	}

	for rowsMembers.Next() {
		var member User

		err := rowsMembers.Scan(&member.Id, &member.Name, &member.Picture)
		if err != nil {
			return chat, err
		}
		chat.Members = append(chat.Members, member)
	}

	for rowsMessages.Next() {
		var message Message

		err := rowsMessages.Scan(&message.Id, &message.Date, &message.Content, &message.Sender.Id, &message.Sender.Name, &message.Sender.Picture, &message.ChatId)
		if err != nil {
			return chat, err
		}

		// Collects all the comments and their senders
		rows, err := db.c.Query("SELECT users.userId, users.userName, users.picture, comments.commentId, comments.content FROM messages JOIN comment_message ON comment_message.messageId = messages.messageId JOIN comments ON comments.commentId = comment_message.commentId JOIN comment_user ON comments.commentId = comment_user.commentId JOIN users ON users.userId = comment_user.userId WHERE messages.messageId = ?", message.Id)
		if err != nil {
			return chat, err
		}
		for rows.Next() {
			var user User
			var comment Comment
			err = rows.Scan(&user.Id, &user.Name, &user.Picture, &comment.Id, &comment.Content)
			if err != nil {
				return chat, err
			}
			var comm_send CommentSender = CommentSender{
				Sender:  user,
				Comment: comment,
			}
			message.CommSenders = append(message.CommSenders, comm_send)
		}
		err = rows.Close()
		if err != nil {
			return chat, err
		}
		chat.Messages = append(chat.Messages, message)
	}
	err = rowsMembers.Close()
	if err != nil {
		return chat, err
	}
	err = rowsMessages.Close()
	if err != nil {
		return chat, err
	}

	return chat, err
}
