package database

func (db *appdbimpl) GetMyConversations(userId string) ([]Chat, error) {

	row, err := db.c.Query("SELECT chats.chatId, chats.chatName FROM chats JOIN chat_user ON chat_user.chatId = chats.chatId WHERE userId = ?", userId)

	if err != nil {
		return nil, err
	}

	var chats []Chat
	defer row.Close()
	for row.Next() {
		var chat Chat
		if row.Scan(&chat.Id, &chat.Name) != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	return chats, err
}
