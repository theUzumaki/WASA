package database

import (
	"errors"
	"strconv"
	"wasatext/service/structs"
)

func (db *appdbimpl) NewChat(userId string, chat structs.Chat) (int, error) {

	// Checks if already exist a private chat with the same users
	if chat.Name == "chat" {
		query := `
			SELECT c.chatId
			FROM chats c
			JOIN chat_user cu1 ON c.chatId = cu1.chatId
			JOIN chat_user cu2 ON c.chatId = cu2.chatId
			WHERE c.chatName = 'chat' AND 
				  ((cu1.userId = ? AND cu2.userId = ?) OR (cu1.userId = ? AND cu2.userId = ?))
			GROUP BY c.chatId
		`
		var chatid int
		err := db.c.QueryRow(query, chat.Members[0].Id, chat.Members[1].Id, chat.Members[1].Id, chat.Members[0].Id).Scan(&chatid)
		if err == nil {
			existingChat, err := db.GetConversation(strconv.Itoa(chatid), userId)
			if err != nil {
				return -1, err
			}
			return existingChat.Id, errors.New("chat already existing")
		} else if err.Error() != "sql: no rows in result set" {
			return -1, err
		} else {
			// No existing chat found, continue with creating a new one
		}
	} else if len(chat.Name) > 32 {
		return 0, errors.New("chat name invalid")
	} else if len(chat.Name) < 3 {
		return 0, errors.New("chat name invalid")
	}

	// Retrieves max id
	row := db.c.QueryRow("SELECT MAX(chatId) FROM chats")
	var stringId string
	var id int
	err := row.Scan(&stringId)
	if stringId == "" {
		id = 0
	} else if err != nil {
		return -1, err
	} else {
		id, err = strconv.Atoi(stringId)
		if err != nil {
			return 0, err
		}
	}
	chat.Id = id + 1

	for i := 0; i < len(chat.Members); i++ {
		for j := 0; j < len(chat.Members); j++ {
			if i != j && chat.Members[i].Id == chat.Members[j].Id {
				return 0, errors.New("duplicate members")
			}
		}
	}

	// Checks if all users exists
	for i := 0; i < len(chat.Members); i++ {
		row := db.c.QueryRow("SELECT * FROM users WHERE userId = ?;", chat.Members[i].Id)
		var user structs.User

		err := row.Scan(&user.Id, &user.Name, &user.Picture)
		if err != nil {
			return 0, err
		}
	}

	_, err = db.c.Exec("INSERT INTO chats VALUES (?, ?, ?);", chat.Id, chat.Name, chat.Picture)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(chat.Members); i++ {
		_, err = db.c.Exec("INSERT INTO chat_user VALUES (?, ?);", chat.Id, chat.Members[i].Id)
		if err != nil {
			return 0, err
		}
	}
	return chat.Id, nil
}
