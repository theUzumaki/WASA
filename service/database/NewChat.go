package database

import (
	"fmt"
	"log"
	"strconv"
)

func (db *appdbimpl) NewChat(userId string, chat Chat) error {

	row := db.c.QueryRow("SELECT MAX(chatId) FROM chats")
	var id int
	row.Scan(&id)
	chat.Id = id + 1

	_, err := db.c.Exec("INSERT INTO chats VALUES (?, ?);", chat.Id, chat.Name)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	for i := 0; i < len(chat.Members); i++ {
		fmt.Println("chatid: " + strconv.Itoa(chat.Id) + " userid: " + strconv.Itoa(chat.Members[i].Id))
		_, err = db.c.Exec("INSERT INTO chat_user VALUES (?, ?);", chat.Id, chat.Members[i].Id)
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
	}
	return nil
}
