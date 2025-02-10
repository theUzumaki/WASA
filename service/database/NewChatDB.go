package database

import (
	"database/sql"
	"errors"
	"strconv"
)

func (db *appdbimpl) NewChat(userId string, chat Chat) (int, error) {

	if len(chat.Members) == 2 {
		rows1, err := db.c.Query("SELECT chatId FROM chats;")
		if err != nil {
			return -1, err
		}
		defer rows1.Close()
		for rows1.Next() {
			var chatid string
			err = rows1.Scan(&chatid)
			if err != nil {
				return -1, err
			} else {
				rows2, err := db.c.Query("SELECT * FROM chat_user WHERE chatId = ?;", chatid)
				if err != nil {
					return -1, err
				}
				defer rows2.Close()
				count := 0
				for rows2.Next() {
					count++
				}
				if count != 2 {
					continue
				}

				rows2, err = db.c.Query("SELECT userId FROM chat_user WHERE chatId = ?;", chatid)
				if err != nil {
					return -1, err
				}
				check := 0
				defer rows2.Close()
				if rows2.Next() {
					var id int
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
						chat, err := db.GetConversation(chatid)
						if err != nil {
							return -1, err
						}
						return chat.Id, errors.New("chat already existing")
					} else if check == 1 && id == chat.Members[1].Id {
						chat, err := db.GetConversation(chatid)
						if err != nil {
							return -1, err
						}
						return chat.Id, errors.New("chat already existing")
					}
				}
			}
		}
	}

	row := db.c.QueryRow("SELECT MAX(chatId) FROM chats")
	var stringId string
	var id int
	err := row.Scan(&stringId)
	if stringId == "" {
		id = 0
	} else if err != nil {
		return -1, err
	}
	chat.Id = id + 1

	for i := 0; i < len(chat.Members); i++ {
		for j := 0; j < len(chat.Members); j++ {
			if i != j && chat.Members[i].Id == chat.Members[j].Id {
				return 0, errors.New("duplicate members")
			}
		}
	}

	for i := 0; i < len(chat.Members); i++ {
		row := db.c.QueryRow("SELECT * FROM users WHERE userId = ?;", chat.Members[i].Id)
		err := row.Scan(nil, nil)
		if err != nil {
			return 0, err
		}
	}

	rows, err := db.c.Query("SELECT chatId FROM chats")
	if err != nil {
		return 0, err
	} else {
		defer rows.Close()
		for rows.Next() {
			var id string
			err = rows.Scan(&id)
			if err != nil {
				return 0, err
			}

			i := 0
			for ; i < len(chat.Members); i++ {
				row = db.c.QueryRow("SELECT * FROM chat_user WHERE userId = ? AND chatId = ?;", userId, id)
				err = row.Scan(nil, nil)
				if errors.Is(err, sql.ErrNoRows) {
					break
				} else if err != nil {
					return 0, err
				}
			}
			if i == len(chat.Members)-1 {
				idInt, err := strconv.Atoi(id)
				if err != nil {
					return 0, err
				}
				return idInt, nil
			}
		}
	}

	_, err = db.c.Exec("INSERT INTO chats VALUES (?, ?);", chat.Id, chat.Name)
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
