package database

import (
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

		row := db.c.QueryRow(`SELECT *  
		FROM chats
		WHERE chatId = ?`, chat.Id)

		if err := row.Scan(&chat.Id, &chat.Name, &chat.Picture); err != nil {
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
			`SELECT	chat_message.messageId, messages.date, messages.content, users.userId, users.userName, users.picture, chats.chatId 
			FROM chats
			JOIN chat_message ON chat_message.chatId = chats.chatId
			JOIN messages ON chat_message.messageId = messages.messageId
			JOIN message_user ON messages.messageId = message_user.messageId
			JOIN users ON message_user.userId = users.userId
			WHERE chats.chatId = ?`, chat.Id)
		if err != nil {
			return nil, err
		}

		for rowsMembers.Next() {
			var member structs.User
			if err := rowsMembers.Err(); err != nil {
				return nil, err
			}

			err := rowsMembers.Scan(&member.Id, &member.Name, &member.Picture)
			if err != nil {
				return nil, err
			}
			chat.Members = append(chat.Members, member)
		}

		for rowsMessages.Next() {
			var message structs.Message
			if err := rowsMessages.Err(); err != nil {
				return nil, err
			}

			err := rowsMessages.Scan(&message.Id, &message.Date, &message.Content, &message.Sender.Id, &message.Sender.Name, &message.Sender.Picture, &message.ChatId)
			if err != nil {
				return nil, err
			}

			// Collects all the comments and their senders
			rows, err := db.c.Query("SELECT users.userId, users.userName, users.picture, comments.commentId, comments.content FROM messages JOIN comment_message ON comment_message.messageId = messages.messageId JOIN comments ON comments.commentId = comment_message.commentId JOIN comment_user ON comments.commentId = comment_user.commentId JOIN users ON users.userId = comment_user.userId WHERE messages.messageId = ?", message.Id)
			if err != nil {
				return nil, err
			}
			for rows.Next() {
				var user structs.User
				var comment structs.Comment
				if rows.Err() != nil {
					return nil, err
				}
				err = rows.Scan(&user.Id, &user.Name, &user.Picture, &comment.Id, &comment.Content)
				if err != nil {
					return nil, err
				}
				var comm_send structs.CommentSender = structs.CommentSender{
					Sender:  user,
					Comment: comment,
				}
				message.CommSenders = append(message.CommSenders, comm_send)
			}
			err = rows.Close()
			if err != nil {
				return nil, err
			}
			chat.Messages = append(chat.Messages, message)
		}
		err = rowsMembers.Close()
		if err != nil {
			return nil, err
		}
		err = rowsMessages.Close()
		if err != nil {
			return nil, err
		}

		chats = append(chats, chat)
	}
	err = rowsChat.Close()
	if err != nil {
		return chats, err
	}

	return chats, err
}
