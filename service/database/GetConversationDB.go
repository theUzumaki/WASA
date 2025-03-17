package database

import (
	"database/sql"
	"log"
	"strconv"
	"wasatext/service/structs"
)

func (db *appdbimpl) GetConversation(chatId string, userId string) (structs.Chat, error) {
	var chat structs.Chat

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
		`SELECT	chat_message.messageId, messages.date, messages.content, messages.checkmark, users.userId, users.userName, users.picture, chats.chatId 
		FROM chats
		JOIN chat_message ON chat_message.chatId = chats.chatId
		JOIN messages ON chat_message.messageId = messages.messageId
		JOIN message_user ON messages.messageId = message_user.messageId
		JOIN users ON message_user.userId = users.userId
		WHERE chats.chatId = ?`, chat.Id)
	if err != nil {
		return chat, err
	}

	var num_members int
	for rowsMembers.Next() {
		var member structs.User
		if err := rowsMembers.Err(); err != nil {
			return chat, err
		}

		err := rowsMembers.Scan(&member.Id, &member.Name, &member.Picture)
		if err != nil {
			return chat, err
		}
		chat.Members = append(chat.Members, member)
		num_members++
	}
	err = rowsMembers.Close()
	if err != nil {
		return chat, err
	}

	var msgs_id1 []string
	var msgs_id2 []string
	for rowsMessages.Next() {
		var message structs.Message
		if err := rowsMessages.Err(); err != nil {
			return chat, err
		}

		err := rowsMessages.Scan(&message.Id, &message.Date, &message.Content, &message.Checkmark, &message.Sender.Id, &message.Sender.Name, &message.Sender.Picture, &message.ChatId)
		if err != nil {
			return chat, err
		}

		// Collects all the comments and their senders
		rows, err := db.c.Query("SELECT users.userId, users.userName, users.picture, comments.commentId, comments.content FROM messages JOIN comment_message ON comment_message.messageId = messages.messageId JOIN comments ON comments.commentId = comment_message.commentId JOIN comment_user ON comments.commentId = comment_user.commentId JOIN users ON users.userId = comment_user.userId WHERE messages.messageId = ?", message.Id)
		if err != nil {
			return chat, err
		}
		for rows.Next() {
			var user structs.User
			var comment structs.Comment
			if rows.Err() != nil {
				return chat, err
			}
			err = rows.Scan(&user.Id, &user.Name, &user.Picture, &comment.Id, &comment.Content)
			if err != nil {
				return chat, err
			}
			var comm_send structs.CommentSender = structs.CommentSender{
				Sender:  user,
				Comment: comment,
			}
			message.CommSenders = append(message.CommSenders, comm_send)
		}
		err = rows.Close()
		if err != nil {
			return chat, err
		}

		var messid string
		var usid string
		err = db.c.QueryRow("SELECT * FROM message_viewer WHERE messageId = ? AND userId = ? LIMIT 1", message.Id, userId).Scan(&messid, &usid)
		if err != nil && err != sql.ErrNoRows {
			return chat, err
		} else if err == sql.ErrNoRows {
			msgs_id1 = append(msgs_id1, strconv.Itoa(message.Id))
		}

		var viewers int
		row = db.c.QueryRow("SELECT COUNT(messageId) FROM message_viewer WHERE messageId = ? AND userId", message.Id, userId)
		err = row.Scan(&viewers)
		if err != nil {
			return chat, err
		}
		if viewers == num_members {
			msgs_id2 = append(msgs_id2, strconv.Itoa(message.Id))
			message.Checkmark = true
		}

		chat.Messages = append(chat.Messages, message)
	}
	err = rowsMessages.Close()
	if err != nil {
		return chat, err
	}

	for _, id := range msgs_id1 {
		_, err = db.c.Exec("INSERT INTO message_viewer VALUES (?, ?)", id, userId)
		if err != nil {
			log.Default().Println("Error inserting into message_viewer: ", err)
			return chat, err
		}
	}

	for _, id := range msgs_id2 {
		_, err = db.c.Exec("UPDATE messages SET checkmark = true WHERE messageId = ?", id)
		if err != nil {
			log.Default().Println("Error updating checkmark: ", err)
		}
	}

	return chat, err
}
