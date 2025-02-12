package database

func (db *appdbimpl) CommentMessage(messageid string, comment string) error {

	var message Message
	row := db.c.QueryRow("SELECT messageId FROM messages WHERE messageId = ?", messageid)
	if err := row.Scan(&message.Id); err != nil {
		return err
	}

	_, err := db.c.Exec("UPDATE messages SET comment = ? WHERE messageId = ?", comment, messageid)
	if err != nil {
		return err
	}

	return err
}
