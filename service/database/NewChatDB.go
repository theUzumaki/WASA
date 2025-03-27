package database

import (
	"errors"
	"strconv"
	"wasatext/service/structs"
)

func (db *appdbimpl) NewChat(userId string, chat structs.Chat) (int, error) {

	// Checks if already exist a private chat with the same users
	if chat.Name == "chat" {
		rows1, err := db.c.Query("SELECT chatId FROM chats;")
		if err != nil {
			return -1, err
		}
		for rows1.Next() {
			var chatid string
			if rows1.Err() != nil {
				return -1, err
			}

			err = rows1.Scan(&chatid)
			if err != nil {
				return -1, err
			} else {
				row := db.c.QueryRow("SELECT chatName FROM chats WHERE chatId = ?;", chatid)
				var name string
				err := row.Scan(&name)
				if err != nil {
					return -1, err
				}

				if name != "chat" {
					continue
				}

				rows2, err := db.c.Query("SELECT userId FROM chat_user WHERE chatId = ?;", chatid)
				if err != nil {
					return -1, err
				}
				check := 0
				if rows2.Next() {
					var id int
					if rows2.Err() != nil {
						return -1, err
					}
					err := rows2.Scan(&id)
					if err != nil {
						return -1, err
					}
					if id == chat.Members[0].Id {
						check = 1
					} else if id == chat.Members[1].Id {
						check = 2
					}
				}
				if check != 0 && rows2.Next() {
					var id int
					err := rows2.Scan(&id)
					if err != nil {
						return -1, err
					}
					if check == 2 && id == chat.Members[0].Id {
						chat, err := db.GetConversation(chatid, userId)
						if err != nil {
							return -1, err
						}
						return chat.Id, errors.New("chat already existing")
					} else if check == 1 && id == chat.Members[1].Id {
						chat, err := db.GetConversation(chatid, userId)
						if err != nil {
							return -1, err
						}
						return chat.Id, errors.New("chat already existing")
					}
				}
				err = rows2.Close()
				if err != nil {
					return 0, err
				}
			}
			err = rows1.Close()
			if err != nil {
				return 0, err
			}

		}
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
