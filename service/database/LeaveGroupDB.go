package database

func (db *appdbimpl) LeaveGroup(chatId string, userId string) error {

	_, err := db.c.Exec("DELETE FROM chat_user WHERE chatId = ? AND userId = ?", chatId, userId)
	if err != nil {
		return err
	}

	rows, err := db.c.Query("SELECT chatId FROM chat_user WHERE chatId = ?", chatId)
	if err != nil {
		return err
	}
	defer rows.Close()

	var i int = 0
	for rows.Next() {
		i++
	}
	if err := rows.Err(); err != nil {
		return err
	}

	if i < 1 {
		_, err := db.c.Exec("DELETE FROM chats WHERE chatId = ?", chatId)
		if err != nil {
			return err
		}
	}

	return err
}
