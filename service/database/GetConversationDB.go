package database

import (
	"strconv"
	"wasatext/service/structs"
)

func (db *appdbimpl) GetConversation(chatId string, userId string) (structs.Chat, error) {
	var chat structs.Chat

	// Fetch chat details
	err := db.c.QueryRow(`SELECT chatId, chatName, picture FROM chats WHERE chatId = ?`, chatId).
		Scan(&chat.Id, &chat.Name, &chat.Picture)
	if err != nil {
		return chat, err
	}

	// Fetch chat members
	rowsMembers, err := db.c.Query(
		`SELECT users.userId, users.userName, users.picture
		FROM chat_user
		JOIN users ON chat_user.userId = users.userId
		WHERE chat_user.chatId = ?`, chat.Id)
	if err != nil {
		return chat, err
	}
	defer func() { err = rowsMembers.Close() }()

	for rowsMembers.Next() {
		if rowsMembers.Err() != nil {
			return chat, rowsMembers.Err()
		}

		var member structs.User
		if err := rowsMembers.Scan(&member.Id, &member.Name, &member.Picture); err != nil {
			return chat, err
		}
		chat.Members = append(chat.Members, member)
	}

	// Fetch messages and related data in a single query
	rowsMessages, err := db.c.Query(
		`SELECT 
			messages.messageId, messages.date, messages.content, messages.checkmark, 
			users.userId, users.userName, users.picture, 
			IFNULL(reply.replyId, -1) AS replyId,
			(SELECT COUNT(*) FROM message_viewer WHERE messageId = messages.messageId) AS viewers
		FROM chat_message
		JOIN messages ON chat_message.messageId = messages.messageId
		JOIN message_user ON messages.messageId = message_user.messageId
		JOIN users ON message_user.userId = users.userId
		LEFT JOIN message_reply AS reply ON messages.messageId = reply.messageId
		WHERE chat_message.chatId = ?`, chat.Id)
	if err != nil {
		return chat, err
	}
	defer func() { err = rowsMessages.Close() }()

	var msgsToMarkViewed []string
	var msgsToUpdateCheckmark []string

	for rowsMessages.Next() {
		if rowsMessages.Err() != nil {
			return chat, rowsMessages.Err()
		}

		var message structs.Message
		var viewers int
		if err := rowsMessages.Scan(&message.Id, &message.Date, &message.Content, &message.Checkmark,
			&message.Sender.Id, &message.Sender.Name, &message.Sender.Picture, &message.ReplyId, &viewers); err != nil {
			return chat, err
		}

		// Fetch comments for the message
		rowsComments, err := db.c.Query(
			`SELECT users.userId, users.userName, users.picture, comments.commentId, comments.content
			FROM comment_message
			JOIN comments ON comment_message.commentId = comments.commentId
			JOIN comment_user ON comments.commentId = comment_user.commentId
			JOIN users ON comment_user.userId = users.userId
			WHERE comment_message.messageId = ?`, message.Id)
		if err != nil {
			return chat, err
		}
		defer func() { err = rowsComments.Close() }()

		for rowsComments.Next() {
			if rowsComments.Err() != nil {
				return chat, rowsComments.Err()
			}
			
			var user structs.User
			var comment structs.Comment
			if err := rowsComments.Scan(&user.Id, &user.Name, &user.Picture, &comment.Id, &comment.Content); err != nil {
				return chat, err
			}
			message.CommSenders = append(message.CommSenders, structs.CommentSender{
				Sender:  user,
				Comment: comment,
			})
		}

		// Check if the message is viewed by the user
		var viewerExists bool
		err = db.c.QueryRow(
			`SELECT EXISTS(SELECT 1 FROM message_viewer WHERE messageId = ? AND userId = ?)`,
			message.Id, userId).Scan(&viewerExists)
		if err != nil {
			return chat, err
		}
		if !viewerExists {
			msgsToMarkViewed = append(msgsToMarkViewed, strconv.Itoa(message.Id))
		}

		// Update checkmark if all members have viewed the message
		if viewers == len(chat.Members) && !message.Checkmark {
			msgsToUpdateCheckmark = append(msgsToUpdateCheckmark, strconv.Itoa(message.Id))
			message.Checkmark = true
		}

		chat.Messages = append(chat.Messages, message)
	}

	// Batch insert message viewers
	if len(msgsToMarkViewed) > 0 {
		query := "INSERT INTO message_viewer (messageId, userId) VALUES "
		args := []interface{}{}
		for _, id := range msgsToMarkViewed {
			query += "(?, ?),"
			args = append(args, id, userId)
		}
		query = query[:len(query)-1] // Remove trailing comma
		if _, err := db.c.Exec(query, args...); err != nil {
			return chat, err
		}
	}

	// Batch update message checkmarks
	if len(msgsToUpdateCheckmark) > 0 {
		query := "UPDATE messages SET checkmark = true WHERE messageId IN ("
		args := []interface{}{}
		for _, id := range msgsToUpdateCheckmark {
			query += "?,"
			args = append(args, id)
		}
		query = query[:len(query)-1] + ")" // Remove trailing comma and close parenthesis
		if _, err := db.c.Exec(query, args...); err != nil {
			return chat, err
		}
	}

	return chat, nil
}
