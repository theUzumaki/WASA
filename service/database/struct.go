package database

import "time"

type User struct {
	userId int `json:"id"`
}

type Message struct {
	messageId int `json:"mess_id"`
	chatId int `json:"chat_id"`
	sender User `json:"sender"`
	date time.Time `json:"date"`
	content string `json:"content"`
	comment string `json:"comment"`
}

type Group struct {
	groupId int `json:"id"`
}

type Chat struct {
	chatId int `json:"id_of_chat"`
	chatName string `json:"chat_name"`
	members []User `json:"members"`
	messages []Message `json:"messages"`
}