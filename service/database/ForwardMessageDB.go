package database

func (db *appdbimpl) ForwardMessage(userid string, messageid string, chatid string) error {

	row := db.c.QueryRow("SELECT MAX(messageId) FROM messages")
	var newid int
	err := row.Scan(&newid)
	if err != nil {
		return err
	}

	row = db.c.QueryRow("SELECT * FROM users WHERE userId = ?;", userid)
	err = row.Scan(nil, nil)
	if err != nil {
		return err
	}

	var message Message
	// Collects all the comments and their senders and stores them
	/*
		rows, err := db.c.Query("SELECT users.userId, users.userName, users.picture, comments.commentId, comments.content FROM messages JOIN comment_message ON comment_message.messageId = messages.messageId JOIN comments ON comments.commentId = comment_message.commentId JOIN comment_user ON comments.commentId = comment_user.commentId JOIN users ON users.userId = comment_user.userId WHERE messages.messageId = ?", messageid)
		if err != nil {
			return err
		}
		for rows.Next() {
			var user User
			var comment Comment
			err = rows.Scan(&user.Id, &user.Name, &user.Picture, &comment.Id, &comment.Content)
			if err != nil {
				return err
			}
			var comm_send CommentSender = CommentSender{
				Sender:  user,
				Comment: comment,
			}
			message.CommSenders = append(message.CommSenders, comm_send)
		}
		err = rows.Close()
		if err != nil {
			return err
		}
	*/

	row = db.c.QueryRow("SELECT * FROM messages WHERE messageId = ?;", messageid)
	err = row.Scan(&message.Id, &message.Date, &message.Content)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("INSERT INTO messages VALUES (?,?,?)", newid+1, message.Date, message.Content)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("INSERT INTO chat_message VALUES (?,?)", chatid, newid+1)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("INSERT INTO message_user VALUES (?,?)", newid+1, userid)
	if err != nil {
		return err
	}

	/*
		for _, commSender := range message.CommSenders {
			commentId := commSender.Comment.Id
			_, err = db.c.Exec("INSERT INTO comment_message VALUES (?,?)", commentId, newid+1)
			if err != nil {
				return err
			}
		}
	*/

	return err
}
